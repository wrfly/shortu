package routes

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wrfly/testing-kit/util/tokenbucket"

	"github.com/wrfly/yasuser/config"
	"github.com/wrfly/yasuser/filter"
	"github.com/wrfly/yasuser/routes/asset"
	s "github.com/wrfly/yasuser/shortener"
	"github.com/wrfly/yasuser/types"
)

const (
	maxPasswdLength = 60
	maxCustomLength = 60
)

var (
	validCustomURI *regexp.Regexp
	indexTemplate  *template.Template
)

func init() {
	validURI, err := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9_+-]+$")
	if err != nil {
		panic(err)
	}
	validCustomURI = validURI

	a, err := asset.Find("/index.html")
	if err != nil {
		panic(err)
	}
	indexTemplate = a.Template()

}

type server struct {
	domain string
	gaID   string
	limit  int64

	handler s.Shortener
	fileMap map[string]bool
	filter  filter.Filter

	host string
	tb   map[string]tokenbucket.Bucket
}

func newServer(conf config.SrvConfig,
	shortener s.Shortener, filter filter.Filter) server {

	u, err := url.Parse(conf.Domain)
	if err != nil {
		panic(err)
	}

	srv := server{
		domain:  conf.Domain,
		host:    u.Host,
		handler: shortener,
		gaID:    conf.GAID,
		limit:   conf.Limit,
		fileMap: make(map[string]bool),
		tb:      make(map[string]tokenbucket.Bucket, 0),
		filter:  filter,
	}
	for _, a := range asset.List() {
		srv.fileMap[a.Name()] = true
	}

	return srv
}

func (s *server) handleIndex() gin.HandlerFunc {
	curlUA := regexp.MustCompile("curl*")

	return func(c *gin.Context) {
		if matched := curlUA.MatchString(c.Request.UserAgent()); matched {
			// query from curl
			c.String(200, fmt.Sprintf("curl %s -d \"%s\"",
				s.domain, "http://longlonglong.com/long/long/long?a=1&b=2"))
		} else {
			// visit from a web browser
			indexTemplate.Execute(c.Writer, map[string]string{
				"domain": s.domain,
				"gaID":   s.gaID,
			})
		}
	}
}

func (s *server) handleURI() gin.HandlerFunc {

	return func(c *gin.Context) {
		URI := c.Param("URI")

		if URI == "" {
			c.String(404, fmt.Sprintln("not found"))
			return
		}
		if s.fileMap[c.Request.RequestURI] {
			asset.Handler(c.Writer, c.Request)
			return
		}

		// handle shortURL
		shortURL, err := s.handler.Restore(URI)
		switch err {
		case types.ErrURLExpired:
			c.String(http.StatusForbidden, "URL expired")
		case types.ErrNotFound:
			c.String(http.StatusNotFound, fmt.Sprintln(err))
		case nil:
			if err := s.invalidURL(shortURL.Ori); err != nil {
				c.String(http.StatusBadGateway, fmt.Sprintln(err))
				return
			}
			c.Redirect(http.StatusTemporaryRedirect, shortURL.Ori)
		default:
			c.String(http.StatusInternalServerError, fmt.Sprintln(err))
		}
	}
}

func (s *server) handleLongURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		// rate limit
		IP := c.ClientIP()
		if tb, ok := s.tb[IP]; !ok {
			s.tb[IP] = tokenbucket.New(s.limit, time.Second)
		} else {
			if !tb.TakeOne() {
				badRequest(c, fmt.Errorf("rate exceeded"))
				return
			}
		}

		buf := urlBufferPool.Get().([]byte)
		defer urlBufferPool.Put(buf)
		n, err := c.Request.Body.Read(buf)
		if err != io.EOF && err != nil {
			badRequest(c, err)
			return
		}
		if n > MAX_URL_LENGTH {
			badRequest(c, types.ErrURLTooLong)
			return
		}

		longURL := fmt.Sprintf("%s", buf[:n])
		if err := s.invalidURL(longURL); err != nil {
			badRequest(c, err)
			return
		}

		opts, err := generateOptions(c.Request.Header)
		if err != nil {
			badRequest(c, err)
			return
		}
		shortURL, err := s.handler.Shorten(longURL, opts)
		if err != nil {
			badRequest(c, err)
			return
		}

		c.String(200, fmt.Sprintf("%s/%s\n", s.domain, shortURL.Short))
	}
}

func badRequest(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, fmt.Sprintln(err.Error()))
}

func generateOptions(h http.Header) (*types.ShortOptions, error) {
	var (
		duration time.Duration = -1
		err      error
	)

	customURI := h.Get("CUSTOM")
	passWord := h.Get("PASS")
	ttl := h.Get("TTL")
	if ttl != "" {
		duration, err = time.ParseDuration(ttl)
		if err != nil {
			return nil, err
		}
	}

	if len(passWord) > maxPasswdLength {
		return nil, fmt.Errorf("passwd length exceeded, max %d",
			maxPasswdLength)
	}

	if customURI != "" {
		if len(customURI) > maxCustomLength {
			return nil, fmt.Errorf("custom URI length exceeded, max %d",
				maxCustomLength)
		}
		if !validCustomURI.MatchString(customURI) {
			return nil, fmt.Errorf("invalid custom URI, must match %s",
				validCustomURI.String())
		}
	}

	return &types.ShortOptions{
		Custom: customURI,
		TTL:    duration,
		Passwd: passWord,
	}, nil
}

func (s *server) invalidURL(URL string) error {
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}

	if u.Hostname() == s.host {
		return types.ErrSameHost
	}
	switch u.Scheme {
	case "http", "https", "ftp", "tcp":
	default:
		return types.ErrScheme
	}

	return s.filter.OK(u)
}
