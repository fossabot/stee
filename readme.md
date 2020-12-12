> **:warning: Unstable**
This project is in a very very early stage. It is incomplete and untested.  
> I consider this project as a training project so you probably shouldn't use it for anything important but any critics on the code are welcomed.

# Stee - URL shortening and more
[![License Card](https://img.shields.io/github/license/milanrodriguez/stee)](LICENSE)
[![Version Card](https://img.shields.io/github/v/release/milanrodriguez/stee?sort=semver)](https://github.com/milanrodriguez/stee/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/milanrodriguez/stee)](https://goreportcard.com/report/github.com/milanrodriguez/stee)

Simple Self-hosted URL redirection service written in Go

## Features

Currently Stee performs HTTP redirects key -> target.  
It persistently stores redirections in a file on disk.

## Roadmap

You can see the roadmap there: https://github.com/milanrodriguez/stee/projects/1  
  
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

Download the binaries here: https://github.com/milanrodriguez/stee/releases.  
Or compile the sources (you can use the [build script](build)).


## Usage

Execute the binaries. A file named "stee.db" will be created where the program has been invoked (pwd).    
Soon, you'll be able to configure the path, see [this issue](https://github.com/milanrodriguez/stee/issues/9).  
  
This early version implements a very simple (and unsecure) HTTP API.  
To add a redirection, the following will work:  

```http
GET http://host:port/_api/simple/add/[key]/[base64URLEncodedTarget]
```

To delete one:  

```http
GET http://host:port/_api/simple/del/[key]
```

And hitting ```http://host:port/[key]``` will redirect you to your target.  

### Configuration

You can see a complete example of a configuration file [here](stee.yaml).  
Flags and environment variables should be available soon, see [this issue](https://github.com/milanrodriguez/stee/issues/8).  

## Maintainer

[@milanrodriguez](https://github.com/milanrodriguez)

## License

See [LICENSE](LICENSE)
