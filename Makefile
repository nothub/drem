MOD_NAME = $(shell go list -m)
BIN_NAME = $(shell basename $(MOD_NAME))
VERSION  = $(shell git describe --tags --abbrev=0 --match v[0-9]* 2> /dev/null || echo "v0.0.0")
LDFLAGS  = -ldflags="-X '$(MOD_NAME)/version.tag=$(VERSION)'"

out/$(BIN_NAME): $(shell ls go.mod go.sum *.go **/*.go)
	go build $(LDFLAGS) -race -o out/$(BIN_NAME)

.PHONY: release
release: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o out/$(BIN_NAME)
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o out/$(BIN_NAME)-arm

.PHONY: check
check: lint test

.PHONY: lint
lint:
	go vet

.PHONY: test
test:
	go test -v -parallel $(shell grep -c -E "^processor.*[0-9]+" "/proc/cpuinfo") $(MOD_NAME)/...

.PHONY: clean
clean:
	go clean
	go mod tidy
	-rm -rf out
