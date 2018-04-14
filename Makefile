#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GOBIN=$(GOPATH)/bin
export CGO_ENABLED=0
export GOOS=linux

all: lint

lint:
	which gometalinter || (go get -u -v github.com/alecthomas/gometalinter && gometalinter --install)
	gometalinter --disable-all --enable=gofmt $(ROOT)
	gometalinter --disable-all --enable=vet $(ROOT)
	gometalinter --disable-all --enable=vetshadow $(ROOT)
	gometalinter --disable-all --enable=gocyclo --cyclo-over=15 $(ROOT)
	gometalinter --disable-all --enable=golint $(ROOT)
	gometalinter --disable-all --enable=ineffassign $(ROOT)
	gometalinter --disable-all --enable=misspell $(ROOT)

format:
	which goimports || go get -u -v golang.org/x/tools/cmd/goimports
	goimports -w $(ROOT)
	gofmt -s -w $(ROOT)
