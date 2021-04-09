#!/bin/bash
set -e

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/shorter

