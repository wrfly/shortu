# Yasuser

> Yet another self-hosted URL shortener.

*Short* or *Restore* your URL, like <https://git.io/> or <https://goo.gl/>
but under **YOUR** control.

[![Master Build Status](https://travis-ci.org/wrfly/yasuser.svg?branch=master)](https://travis-ci.org/wrfly/yasuser)
[![Go Report Card](https://goreportcard.com/badge/github.com/wrfly/yasuser)](https://goreportcard.com/report/github.com/wrfly/yasuser)
[![license](https://img.shields.io/github/license/wrfly/yasuser.svg)](https://github.com/wrfly/yasuser/blob/master/LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/wrfly/yasuser.svg)](https://hub.docker.com/r/wrfly/yasuser/)

## Run

```sh
docker run --name yasuser -dti \
    -p 8080:8084 \
    -e YASUSER_SHORTENER_STORE_DBPATH=/data/yasuser.db \
    -e YASUSER_SERVER_DOMAIN=your.domain.com \
    -v `pwd`:/data \
    wrfly/yasuser
```

Or use the [docker-compose.yml](./docker-compose.yml).

## Usage

```bash
# short your URL
➜  ~ curl https://u.kfd.me -d "https://kfd.me"
https://u.kfd.me/1C
➜  ~

# restore it
➜  ~ curl http://u.kfd.me/1C
<a href="https://kfd.me">Found</a>.

```

Or just visit the web page:

![index](index.png)

## Features

- [x] it works
- [x] blob database
- [x] length and validate
- [x] memory cache
- [x] redis database
- [ ] customization
- [ ] TTL of URL
- [ ] rate limit
- [ ] management(auth)
  - [ ] remove(domain or keywords)
  - [ ] blacklist(domain or keywords)
  - [ ] whitelist(domain or keywords)
- [ ] statistic
  - [ ] URL status
  - [ ] runtime metrics
- [x] UI index
