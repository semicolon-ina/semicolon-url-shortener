#!/bin/sh
# source .env

go mod tidy && go mod vendor

go run main.go
