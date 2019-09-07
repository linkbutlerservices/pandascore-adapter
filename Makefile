.DEFAULT_GOAL := build
.PHONY: build install docker dockerpush

REPO=linkbutlerservices/pandascore-adapter
LDFLAGS=-ldflags "-X github.com/linkbutlerservices/pandascore-adapter/store.Sha=`git rev-parse HEAD`"

build:
	@go build $(LDFLAGS) -o pandascore-adapter

install:
	@go install $(LDFLAGS)

docker:
	@docker build . -t $(REPO)

dockerpush:
	@docker push $(REPO)