> **:warning: Unstable**
This project is in a very very early stage. It is incomplete and untested.
> I consider this project as a training project so you probably shouldn't use it for anything important but any critics on the code are welcomed.

# Stee - URL shortening and redirection

Simple Self-hosted URL redirection service written in Go

## Features

Currently Stee can basically redirect http clients.  
It stores redirections in a file in a persistent way.

## Roadmap

In this order:

- [ ] Tests
- [ ] HTTPS
- [ ] proper API and auth
- [ ] Extra functionalities (password protection, captcha, TTL, etc)
- [ ] Metrics & analytics
- [ ] web app UI

## Setup

Download the binaries or compile the sources.

## Usage

Execute the binaries. A file named "stee.db" will be created where the program has been invoked (pwd).  
  
This early version implements a very simple (and unsecure) HTTP API.  
To add a redirection, the following will work:  

```http
GET http://host:port/_api/simple/add/[key]/[base64URLEncodedTarget]
```

To delete one:  

```http
GET http://host:port/_api/simple/del/[key]
```

Then, hitting ```http://host:port/[key]``` will redirect you to your target.

## Maintainer

@milanrodriguez

## License

© Milan Rodriguez
