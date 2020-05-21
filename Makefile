GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
mods:
	$(GOMOD) download

build-osx:
	env GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -v ./...
build-linux:
	env GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -v ./...
build-windows:
	env GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -v ./...