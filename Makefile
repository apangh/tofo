# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOINSTALL=$(GOCMD) install
MOCKGEN=$(GOPATH)/bin/mockgen

mod:
	@go mod download

clean:
	$(GOCLEAN)

vet:
	$(GOVET) ./...

test:
	$(GOTEST) -race ./...

test_all:
	$(GOTEST) all

build:
	$(GOBUILD) ./...
