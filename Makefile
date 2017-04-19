PACKAGES:=$$(go list ./... | grep -v /vendor/)

.phony: build test fmt

build:
	mkdir -p ./bin
	go build -o ./bin/metronome .

install:
	@glide install --strip-vendor

update:
	@glide update --strip-vendor

fmt:
	@go fmt $(PACKAGES)

lint:
	@golint ./... | grep -vE "vendor" || printf ""

test:
	@go test ${PACKAGES}
