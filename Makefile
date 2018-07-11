VALIDATE_DEPS = github.com/golang/lint/golint
TEST_DEPS     = github.com/axw/gocov/gocov github.com/AlekSi/gocov-xml
GO_PKGS      := $(shell go list ./... | grep -v "/vendor/")
GO_FILES     := $(shell find . -type f -name "*.go" | grep -v "/vendor/" | grep -v "bindata.go")
GOTOOLS       = github.com/kardianos/govendor github.com/jteeuwen/go-bindata/... \
		github.com/kardianos/govendor \
		gopkg.in/alecthomas/gometalinter.v2 \
		github.com/axw/gocov/gocov \
		github.com/AlekSi/gocov-xml

all: clean deps validate test compile

clean:
	@echo "[ clean ]: removing binaries and coverage file..."
	@rm -rfv bin coverage.xml

validate-deps:
	@echo "[ validate-deps ]: installing validation dependencies..."
	@go get -v $(VALIDATE_DEPS)

validate: validate-deps
	@printf "[ validate ]: running gofmt... "
	@gometalinter.v2 --config=.gometalinter.json ./...

generate:
	@go generate $(GO_PKGS)

tools:
	@echo "[ tools ]: installing compilation dependencies..."
	@go get -v $(GOTOOLS)
	@gometalinter.v2 --install

tools-update:
	@echo "[ tools-update ]: updating tools..."
	@go get -u $(GOTOOLS)
	@gometalinter.v2 --install

compile: tools
	@echo "[ compile ]: building 'nr-integrations-builder'..."
	@go build -o bin/nr-integrations-builder

test-deps: tools
	@echo "[ test-deps ]: installing testing dependencies..."
	@go get -v $(TEST_DEPS)

test: test-deps
	@echo "[ test ]: running unit tests..."
	@gocov test $(GO_PKGS) | gocov-xml > coverage.xml

deps: tools
	@echo "[ install ]: installing..."
	@go get $(GO_PKGS)

.PHONY: all clean validate-deps validate generate tools compile test-deps test deps
