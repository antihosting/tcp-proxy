APP_NAME := port_proxy

BIN := $(APP_NAME)
BIN_MAC := $(BIN)_darwin
BIN_LINUX := $(BIN)_linux
BIN_WINDOWS := $(BIN).exe
TARGET := distr

VERSION := $(shell git describe --tags --always --dirty)
NOW := $(shell date +"%m-%d-%Y")
HASHED_TOKEN := $(shell cat .hashed-token.txt)

build:
	go build -o $(BIN) -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW) -X main.HashedToken=$(HASHED_TOKEN)"

distr: build
	rm -rf $(TARGET)
	mkdir $(TARGET)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(TARGET)/$(BIN_LINUX) -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	gzip $(TARGET)/$(BIN_LINUX)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(TARGET)/$(BIN_MAC) -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	gzip $(TARGET)/$(BIN_MAC)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -o $(TARGET)/$(BIN_WINDOWS) -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	gzip $(TARGET)/$(BIN_WINDOWS)

update:
	go get -u ./...

test:
	go test -v

clean:
	go clean
	rm -f $(BIN)
	rm -f $(BIN_LINUX)
	rm -f $(BIN_WINDOWS)
