package routes

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"regexp"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"

	"github.com/wrfly/yasuser/shortener"
	"github.com/wrfly/yasuser/types"
)

type server struct {
	domain string
	gaID   string

	stener        shortener.Shortener
	indexTemplate *template.Template
	fileMap       map[string]bool

	host string
}

func (s *server) init() {
	fileMap := make(map[string]bool, len(AssetNames()))
	for _, fileName := range AssetNames() {
		fileMap[fileName] = true
	}
	s.fileMap = fileMap

	bs, err := Asset("index.html")
	if err != nil {
		panic(err)
	}
	s.indexTemplate, err = template.New("index").Parse(string(bs))
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(s.domain)
	if err != nil {
		panic(err)
	}
	s.host = u.Host
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
			s.indexTemplate.Execute(c.Writer, map[string]string{
				"domain": s.domain,
				"gaID":   s.gaID,
			})
		}
	}
}

func (s *server) handleURI() gin.HandlerFunc {

	return func(c *gin.Context) {
		URI := c.Param("URI")

		switch {
		case URI == "":
			c.String(404, fmt.Sprintln("not found"))

		case s.fileMap[URI]:
			// handle static files
			http.FileServer(&assetfs.AssetFS{
				Asset:     Asset,
				AssetDir:  AssetDir,
				AssetInfo: AssetInfo,
				Prefix:    "/",
			}).ServeHTTP(c.Writer, c.Request)

		default:
			// handle shortURL
			if realURL := s.stener.Restore(URI); realURL == "" {
				c.String(404, fmt.Sprintln("not found"))
			} else {
				c.Redirect(302, realURL)
			}
		}
	}
}

func (s *server) handleLongURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := urlBufferPool.Get().([]byte)
		defer urlBufferPool.Put(buf)
		n, err := c.Request.Body.Read(buf)
		if err != io.EOF && err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s\n", err))
			return
		}
		if n > MAX_URL_LENGTH {
			c.String(http.StatusBadRequest, fmt.Sprintln("Bad request, URL too long"))
			return
		}

		longURL := fmt.Sprintf("%s", buf[:n])
		if s.invalidURL(longURL) {
			c.String(http.StatusBadRequest, fmt.Sprintln("Invalid URL"))
			return
		}

		var short string
		if customURL := c.Request.Header.Get("CUSTOM"); customURL != "" {
			// custom shorten URL
			if err := s.stener.ShortenWithCustomURL(customURL, longURL); err != nil {
				if err == types.ErrAlreadyExist {
					c.String(http.StatusBadRequest, fmt.Sprintln(err.Error()))
				} else {
					c.String(http.StatusInternalServerError,
						"something bad happend\n")
				}
				return
			}
			short = customURL
		} else { // default
			if short = s.stener.Shorten(longURL); short == "" {
				c.String(http.StatusInternalServerError,
					"something bad happend\n")
				return
			}
		}

		// TODO: shorten URL with TTL

		shortURL := fmt.Sprintf("%s/%s", s.domain, short)
		c.String(200, fmt.Sprintln(shortURL))
	}
}

func (s *server) invalidURL(URL string) bool {
	u, err := url.Parse(URL)
	if err != nil {
		return true
	}

	if u.Host == s.host {
		return true
	}

	switch u.Scheme {
	case "http", "https", "ftp", "tcp":
		return false
	default:
		return true
	}
}
