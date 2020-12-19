#!/bin/sh
go build -mod vendor -o bin/stee -ldflags '-w -s' .
