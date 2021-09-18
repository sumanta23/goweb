GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BUILD_DIR="bin"
OUTPUT=$(BUILD_DIR)/app

all: test vet fmt lint build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status


build:
	mkdir -p $(BUILD_DIR)
	$(GO_BUILD_ENV) go build -v -o $(OUTPUT) .

clean:
	rm -rf $(BUILD_DIR)
