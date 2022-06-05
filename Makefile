REPO_ROOT = $(shell git rev-parse --show-toplevel)
BIN_DIR   = $(REPO_ROOT)/bin

$(BIN_DIR):
	mkdir -p $@

######################################################################

NAME = condition-bitbucket
BIN  = $(BIN_DIR)/$(NAME)

######################################################################

GOLANGCILINT_VERSION = v1.46.2
GOLANGCILINT_NAME    = golangci-lint
GOLANGCILINT_PATH    = $(REPO_ROOT)/bin
GOLANGCILINT_URL     = https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh
GOLANGCILINT         = $(REPO_ROOT)/bin/$(GOLANGCILINT_NAME)

$(GOLANGCILINT):
	cd $(REPO_ROOT); curl -sSfL $(GOLANGCILINT_URL) | sh -s -- -b $(BIN_DIR) $(GOLANGCILINT_VERSION)

######################################################################



######################################################################

.PHONY: clean
clean:
	rm -f $(BIN)

.PHONY: lint
lint: $(GOLANGCILINT)
	$(GOLANGCILINT) run

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -o $(BIN) $(REPO_ROOT)/cmd/condition-bitbucket
