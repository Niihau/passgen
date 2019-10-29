GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=passgen

default:
	@echo "build target is required"
	@exit 2
build:
	time $(GOBUILD) -o $(BINARY_NAME) -v
run: build
	./$(BINARY_NAME)
buildlinux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
