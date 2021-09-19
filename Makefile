ifeq ($(OS),Windows_NT) 
	detected_OS := Windows
else
	detected_OS := $(shell bash -c 'uname 2>/dev/null || echo Unknown')
endif

ifeq ($(detected_OS), Windows)
	TARGET := windows
endif
ifeq ($(detected_OS), Linux)
	TARGET := linux
endif
ifeq ($(detected_OS), Darwin)
	TARGET := darwin
endif

$(info ************  $(TARGET)  ************)
GO=go
GOCOVER=$(GO) tool cover

GO_BUILD_ENV := CGO_ENABLED=0 GOOS=$(TARGET) GOARCH=amd64
BUILD_DIR=bin
DOCKER=docker
OUTPUT=$(BUILD_DIR)/app
COVERAGE=$(BUILD_DIR)/coverage.out

all: test vet fmt lint build

clean:
	rm -rf $(BUILD_DIR)

gc:
	rm -rf coverage
	rm -rf tags

dir:
	mkdir -p $(BUILD_DIR)

tags:
	gotags -R * > tags

test:
	$(GO) test ./...

vet:
	$(GO) vet ./...

fmt:
	$(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	test -z $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	$(GO) list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status


build: clean dir
	$(GO_BUILD_ENV) $(GO) build -o $(OUTPUT) .

docker:
	$(DOCKER) build  -t goapp -f .docker/Dockerfile . 
