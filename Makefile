ALL_BINARIES := build/linux \
	build/arm
GO ?= /home/aidan/go/go1.16.5/bin/go



all: $(ALL_BINARIES)

.PHONY: clean \
	version \
	mod_update \
	test \
	$(ALL_BINARIES)

build/linux:
	(env GOOS=linux GOARCH=arm GOARM=7 $(GO) build -v -mod=vendor)

build/arm:
	(env GOOS=linux GOARCH=arm GOARM=7 $(GO) build -v -mod=vendor)


clean:
	-rm $(ALL_BINARIES)

version:
	$(GO) version

mod_update:
	$(GO) get -u all

test:
	go test -mod=vendor ./...
