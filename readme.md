> **:warning: Unstable**  
This project is in a very very early stage. It is incomplete and untested.  
> I consider this project as a training project so you probably shouldn't use it for anything important but any critics on the code are welcomed.

# Stee - URL shortening and more

[![License Card](https://img.shields.io/github/license/milanrodriguez/stee)](LICENSE)
[![Version Card](https://img.shields.io/github/v/release/milanrodriguez/stee?sort=semver)](https://github.com/milanrodriguez/stee/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/milanrodriguez/stee)](https://goreportcard.com/report/github.com/milanrodriguez/stee)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmilanrodriguez%2Fstee.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmilanrodriguez%2Fstee?ref=badge_shield)

Simple Self-hosted URL redirection service written in Go

## Features

Currently Stee performs HTTP redirects domain.dev/key -> target.  
It provides a simple API to CRUD the redirections.
It persistently stores redirections in a file on disk.

## Roadmap

You can see the roadmap there: <https://github.com/milanrodriguez/stee/projects/1>  
  
What's planned:

- [x] Tests
- [x] File-based configuration (+flags +env)
  - [ ] Flags
  - [ ] Env variables
- [x] HTTPS
- [ ] Autogeneration of links (true URL shortener)
- [ ] proper API and auth
- [ ] Extra functionalities (password protection, captcha, TTL, etc)
- [ ] More datasources (KV are particularly indicated: etcd, redis, or other databases)
- [ ] Metrics & analytics
- [ ] web app UI

## Setup

Download the binaries here: <https://github.com/milanrodriguez/stee/releases>.  
Or compile the sources (you can use the [build script](build)).

## Usage

To launch Stee, you need to give it a configuration file. See [Configuration](#Configuration) for more information.  
  
This early version implements a very simple (and unsecure) HTTP API.  
  
To add a redirection:  

```http
GET http://host:port/_api/simple/add/[key]/[base64URLEncodedTarget]
```

To get one:  

```http
GET http://host:port/_api/simple/get/[key]
```

To delete one:  

```http
GET http://host:port/_api/simple/del/[key]
```

And hitting ```http(s)://host:port/[key]``` will redirect you to your target.  

### Configuration

You can see a complete example of a configuration file [here](stee.yaml).  

Stee will look for a configuration file named "stee.yaml" in the current working directory or /etc/stee/ (in this order).

Flags and environment variables should be available soon, see [this issue](https://github.com/milanrodriguez/stee/issues/8).  

## Maintainer

[@milanrodriguez](https://github.com/milanrodriguez)

## License

See [LICENSE](LICENSE)


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmilanrodriguez%2Fstee.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmilanrodriguez%2Fstee?ref=badge_large)