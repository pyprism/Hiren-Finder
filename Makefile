
.PHONY: all get clean build

GO ?= go

all: build

build: get
	${GO} build -o bin/hiren-finder hiren.go

get:
	dep ensure

clean:
	@rm -rf bin/hiren-finder
