VECHO = @
ROOT := github.com/hatech/backup
IMAGE_NAME := wangfengteng/hateach-k8s-backup
VERSION ?= latest
GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)
GO_FLAGS ?= GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 GO111MODULE=on
GOPATH ?= "$(HOME)/go"
GOROOT ?= "$(shell go env GOROOT)"

#IMAGE_VERSION := ${VERSION}-${COMMIT}
IMAGE_VERSION := latest
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /test)

all: mod compile

fmt:
	find pkg cmd -type f -name "*.go" | xargs gofmt -l -w

doc:
	swag init -g cmd/server/main.go -pd -q

mod:
	go mod tidy
	go mod vendor

compile: server-compile

server-compile: ##GOARCH=amd64 GOOS=linux
	$(GO_FLAGS) go build --mod=vendor \
            -ldflags  \
            "-X '$(ROOT)/pkg/version.Version=$(VERSION)'" \
            -a -o ./bin/app ./cmd/server/


image:
	docker build \
	--build-arg VERSION=$(VERSION) \
	-t $(IMAGE_NAME):$(IMAGE_VERSION) .

test:
	go test -cover -gcflags=-l $(PKGS)



