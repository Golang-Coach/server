#!/usr/bin/env bash
go get github.com/kardianos/govendor

go get -u github.com/swaggo/swag/cmd/swag

go build main.go

