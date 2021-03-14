REPO_ROOT = $(shell git rev-parse --show-toplevel)
BIN_DIR   = $(REPO_ROOT)/bin

$(BIN_DIR):
	mkdir -p $@

######################################################################

NAME = condition-bitbucket
BIN  = $(BIN_DIR)/$(NAME)

######################################################################

GOLANGCILINT_VERSION = v1.38.0
GOLANGCILINT_NAME    = golangci-lint
GOLANGCILINT_PATH    = $(REPO_ROOT)/bin
GOLANGCILINT_URL     = https://install.goreleaser.com/github.com/golangci/golangci-lint.sh
GOLANGCILINT         = $(REPO_ROOT)/bin/$(GOLANGCILINT_NAME)

$(GOLANGCILINT):
	cd $(REPO_ROOT); curl -sfL $(GOLANGCILINT_URL) | sh -s $(GOLANGCILINT_VERSION)

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
