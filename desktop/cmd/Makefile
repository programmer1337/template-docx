##-----------------------------##
MAIN_PATH=../main.go
BUILD_PATH=../build/
BUILD_NAME=counterparties.exe
##-----------------------------##
LDFLAGS=-H=windowsgui

all: build

GOOS=windows
GOARCH=amd64
CGO_ENABLED=1
CC=x86_64-w64-mingw32-gcc

build: 
	@echo "Building for GOOS=$(GOOS), GOARCH=$(GOARCH)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) go build -v -o $(BUILD_PATH)$(BUILD_NAME) -ldflags $(LDFLAGS) $(MAIN_PATH)