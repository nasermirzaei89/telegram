#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GOBIN=$(GOPATH)/bin
export CGO_ENABLED=0
export GOOS=linux
export ENV=development

export GLIDE_HOME=$(HOME)/.glide

export LDFLAGS="-w -s"

all: lint

fetch: glide-install

contributors:
	git log --all --format='%aN <%cE>' | sort -u  > CONTRIBUTORS

######
# Lint
######

check-gometalinter:
	which gometalinter || (go get -u -v github.com/alecthomas/gometalinter && gometalinter --install)

lint: fetch check-gometalinter
	gometalinter --vendor --skip=vendor/ --exclude=vendor \
	--disable-all \
	--enable=gofmt \
	--enable=vet --enable=vetshadow \
	--enable=gocyclo \
	--enable=golint \
	--enable=ineffassign \
	--enable=misspell \
	--deadline=5m \
	--concurrency=1 \
	./...

format:
	which gometalinter || go get -u -v golang.org/x/tools/cmd/goimports
	find $(ROOT)/ -type f -name "*.go" | grep -v $(ROOT)/vendor | xargs --max-args=1 --replace=R goimports -w R
	find $(ROOT)/ -type f -name "*.go" | grep -v $(ROOT)/vendor | xargs --max-args=1 --replace=R gofmt -s -w R

#######
# Glide
#######

check-glide: check-glide
	which glide || curl https://glide.sh/get | sh

check-glide-init:
	@[ -f $(ROOT)/glide.yaml ] || make -f $(ROOT)/Makefile glide-init

glide-init: check-glide
	glide init

glide-update: check-glide check-glide-init
	glide update

glide-install: check-glide check-glide-init
	glide install
