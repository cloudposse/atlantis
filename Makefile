BUILD_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo no-commit-id)
WORKSPACE := $(shell pwd)
PKG := $(shell go list ./... | grep -v e2e | grep -v vendor | grep -v static | grep -v mocks | grep -v testing)
PKG_COMMAS := $(shell go list ./... | grep -v e2e | grep -v vendor | grep -v static | grep -v mocks | grep -v testing | tr '\n' ',')
IMAGE_NAME := runatlantis/atlantis

SHELL = /bin/bash
PATH:=$(PATH):$(GOPATH)/bin

-include $(shell curl -sSL -o .build-harness "https://git.io/build-harness"; echo .build-harness)

.PHONY: test

.DEFAULT_GOAL := help
help: ## List targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

id: ## Output BUILD_ID being used
	@echo $(BUILD_ID)

debug: ## Output internal make variables
	@echo BUILD_ID = $(BUILD_ID)
	@echo IMAGE_NAME = $(IMAGE_NAME)
	@echo WORKSPACE = $(WORKSPACE)
	@echo PKG = $(PKG)

deps: ## Download dependencies
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build-service: ## Build the main Go service
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o atlantis .

go-generate: ## Run go generate in all packages
	go generate $(PKG)

#regen-mocks: ## Delete all mocks and matchers and then run go generate to regen them. This doesn't work anymore.
#find . -type f | grep mocks/mock_ | grep -v vendor | xargs rm
#find . -type f | grep mocks/matchers | grep -v vendor | xargs rm
#@# not using $(PKG) here because that it includes directories that have now
#@# been deleted, causing go generate to fail.
#echo "this doesn't work anymore: go generate \$\$(go list ./... | grep -v e2e | grep -v vendor | grep -v static)"

test: ## Run tests
	@go test -short $(PKG)

test-all: ## Run tests including integration
	@go test $(PKG)

test-coverage:
	@mkdir -p .cover
	@go test -coverpkg $(PKG_COMMAS) -coverprofile .cover/cover.out $(PKG)

test-coverage-html:
	@mkdir -p .cover
	@go test -coverpkg $(PKG_COMMAS) -coverprofile .cover/cover.out $(PKG)
	go tool cover -html .cover/cover.out

dist: ## Package up everything in static/ using go-bindata-assetfs so it can be served by a single binary
	rm -f server/static/bindata_assetfs.go && go-bindata-assetfs -pkg static -prefix server server/static/... && mv bindata_assetfs.go server/static

release: ## Create packages for a release
	./scripts/binary-release.sh

fmt: ## Run goimports (which also formats)
	goimports -w $$(find . -type f -name '*.go' ! -path "./vendor/*" ! -path "./server/static/bindata_assetfs.go" ! -path "**/mocks/*")

gometalint: ## Run every linter ever
	# gotype and gotypex are disabled because they don't pass on CI and https://github.com/alecthomas/gometalinter/issues/206
	# maligned is disabled because I'd rather have alphabetical struct fields than save a few bytes
	# gocyclo is temporarily disabled because we don't pass it right now
	# golint is temporarily disabled because we need to add comments everywhere first
	# CGO_ENABLED=0 is attempted workaround for https://github.com/alecthomas/gometalinter/issues/149
	CGO_ENABLED=0 gometalinter --config=.gometalinter.json ./...

gometalint-install: ## Install gometalint
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

check-gometalint: gometalint-install gometalint

check-fmt: ## Fail if not formatted
	go get golang.org/x/tools/cmd/goimports
	goimports -d $$(find . -type f -name '*.go' ! -path "./vendor/*" ! -path "./server/static/bindata_assetfs.go" ! -path "**/mocks/*")

end-to-end-deps: ## Install e2e dependencies
	./scripts/e2e-deps.sh

end-to-end-tests: ## Run e2e tests
	./scripts/e2e.sh

website-dev:
	yarn website:dev

go/get:
	go get

go/build:
	CGO_ENABLED=0 go build -v -o "./dist/bin/atlantis" *.go
