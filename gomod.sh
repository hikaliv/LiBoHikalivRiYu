#! /bin/sh

go mod download
go mod tidy
go mod vendor
go mod verify