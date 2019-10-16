#!/bin/bash

if [[ -z "$1" ]]; then
  echo "Missing output pattern parameter"
  exit 1
fi

go get -u github.com/mitchellh/gox
go get -u github.com/golang/dep/cmd/dep

dep ensure

gox -osarch="windows/386 windows/amd64 freebsd/arm netbsd/386 netbsd/amd64 netbsd/arm linux/s390x linux/arm darwin/386 darwin/amd64 linux/386 linux/amd64 freebsd/amd64 freebsd/386 openbsd/386 openbsd/amd64" \
  -output "$1/atlantis_{{.OS}}_{{.Arch}}"
