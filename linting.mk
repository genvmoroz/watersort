### linting options below
### lint: for local environment
### lint-docker-nexus: simulates the CI outcome on a local environment

TMP_DIR = /tmp/

LINT_VERSION = 1.44.0

LINT_DIR = $(TMP_DIR)/golangci-lint/$(LINT_VERSION)
LINT_BIN = $(LINT_DIR)/golangci-lint

CMD = golangci-lint run --allow-parallel-runners
LINT_ARCHIVE = golangci-lint-$(LINT_VERSION)-$(OSNAME)-amd64.tar.gz
LINT_ARCHIVE_DEST = $(TMP_DIR)/$(LINT_ARCHIVE)

# Run this on local machine.
# It downloads a version of golangci-lint and execute it locally.
.PHONY: lint
lint: $(LINT_BIN)
	$(LINT_DIR)/$(CMD)

# install a local golangci-lint if not found.
$(LINT_BIN):
	curl -L --create-dirs \
		https://github.com/golangci/golangci-lint/releases/download/v$(LINT_VERSION)/$(LINT_ARCHIVE) \
		--output $(LINT_ARCHIVE_DEST)
	mkdir -p $(LINT_DIR)/
	tar -xf $(LINT_ARCHIVE_DEST) --strip-components=1 -C $(LINT_DIR)/
	chmod +x $(LINT_BIN)
	rm -f $(LINT_ARCHIVE_DEST)
