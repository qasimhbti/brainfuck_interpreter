export GO_BUILD=go build
export GO_GET=go get
export GO_TEST=go test

all: build test clean

build:
	$(GO_BUILD) -v -i -o build/brainfuck_interpreter ./interpreter

test:
	$(GO_TEST) -v  ./interpreter
	$(GO_TEST) -v  ./utils

clean:
	rm -rf build

.PHONY: build test clean