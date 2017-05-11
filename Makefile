VALIDATE_DEPS = github.com/golang/lint/golint
TEST_DEPS     = github.com/axw/gocov/gocov github.com/AlekSi/gocov-xml
DEPS          = github.com/kardianos/govendor github.com/jteeuwen/go-bindata/...
GO_PKGS      := $(shell go list ./... | grep -v "/vendor/")
GO_FILES     := $(shell find . -type f -name "*.go" | grep -v "/vendor/" | grep -v "bindata.go")

all: clean validate test compile

clean:
	@echo "[ clean ]: removing binaries and coverage file..."
	@rm -rfv bin coverage.xml

validate-deps:
	@echo "[ validate-deps ]: installing validation dependencies..."
	@go get -v $(VALIDATE_DEPS)

validate: validate-deps
	@printf "[ validate ]: running gofmt... "
	@OUTPUT="$(shell gofmt -l $(GO_FILES))" ;\
	if [ -z "$$OUTPUT" ]; then \
		echo "passed." ;\
	else \
		echo "failed. Incorrect syntax in the following files:" ;\
		echo "$$OUTPUT" ;\
		exit 1 ;\
	fi
	@printf "[ validate ]: running golint... "
	@OUTPUT="$(shell golint $(GO_PKGS) | grep -v bindata)" ;\
	if [ -z "$$OUTPUT" ]; then \
		echo "passed." ;\
	else \
		echo "failed. Issues found:" ;\
		echo "$$OUTPUT" ;\
		exit 1 ;\
	fi
	@printf "[ validate ]: running go vet... "
	@OUTPUT="$(shell go vet $(GO_PKGS))" ;\
	if [ -z "$$OUTPUT" ]; then \
		echo "passed." ;\
	else \
		echo "failed. Issues found:" ;\
		echo "$$OUTPUT" ;\
		exit 1;\
	fi

generate:
	@go generate $(GO_PKGS)

compile-deps:
	@echo "[ compile-deps ]: installing compilation dependencies..."
	@go get -v $(DEPS)

compile: compile-deps
	@echo "[ compile ]: building 'nr-integrations-builder'..."
	@go build -o bin/nr-integrations-builder

test-deps: compile-deps
	@echo "[ test-deps ]: installing testing dependencies..."
	@go get -v $(TEST_DEPS)

test: test-deps
	@echo "[ test ]: running unit tests..."
	@gocov test $(GO_PKGS) | gocov-xml > coverage.xml

install:
	@echo "TODO"

.PHONY: all clean validate-deps validate generate compile-deps compile test-deps test install
