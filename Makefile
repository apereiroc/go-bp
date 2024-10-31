APP_NAME := go-bp
VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
BUILD_DIR := target/build
RELEASE_DIR := target/release
TARGETS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64
GITHUB_OWNER := your-github-username
GITHUB_REPO := $(APP_NAME)
GITHUB_TOKEN := $(shell echo $$GITHUB_TOKEN)
LDFLAGS := "-s -w"  # Strip debugging information to reduce binary size


# Compile binary for each specified platform with optimization flags
define build_target
	$(info Building for GOOS=$(1), GOARCH=$(2))
	GOOS=$(1) GOARCH=$(2) go build -ldflags=$(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-$(1)-$(2) main.go
endef

all: build

# Create directories if they do not exist
$(BUILD_DIR) $(RELEASE_DIR):
	mkdir -p $(BUILD_DIR) $(RELEASE_DIR)

.PHONY: build
build:
	$(info Building optimized binaries for targets: $(TARGETS))
	$(foreach target,$(TARGETS),$(eval OS := $(word 1,$(subst /, ,$(target)))) \
		$(eval ARCH := $(word 2,$(subst /, ,$(target)))) \
		$(call build_target,$(OS),$(ARCH)))

# Package binaries for release
.PHONY: package
package: build $(RELEASE_DIR)
	@echo "Packaging binaries for release"
	$(foreach target,$(TARGETS), \
		$(eval OS := $(word 1,$(subst /, ,$(target)))) \
		$(eval ARCH := $(word 2,$(subst /, ,$(target)))) \
		zip -j $(RELEASE_DIR)/$(APP_NAME)-$(VERSION)-$(OS)-$(ARCH).zip $(BUILD_DIR)/$(APP_NAME)-$(OS)-$(ARCH);)

# Publish release on GitHub
.PHONY: publish-release
publish-release: package
	$(info Publishing release $(VERSION))
	gh release create $(VERSION) $(RELEASE_DIR)/* --title "$(VERSION)" --notes "Release $(VERSION)"

# Run tests
.PHONY: test
test:
	$(info Running tests)
	go test ./...

# Run tests with coverage and generate coverage report
.PHONY: test-coverage
test-coverage:
	$(info Runing tests with coverage)
	go test ./... -coverprofile=coverage.out
	$(info Generating coverage report)
	go tool cover -html=coverage.out -o coverage.html
	$(info Coverage report saved as coverage.html)

# Run linter (requires golangci-lint or golint installed)
.PHONY: lint
lint:
	$(info Running linter)
	golangci-lint run ./...

# Clean build and release directories
.PHONY: clean
clean:
	$(info Cleaning up build and release directories)
	rm -rf $(BUILD_DIR) $(RELEASE_DIR) coverage.out coverage.html

# Display available targets
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build          - Build optimized binaries for all targets"
	@echo "  package        - Package binaries for release"
	@echo "  publish-release - Publish release on GitHub (requires GitHub CLI)"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  lint           - Run linter on the codebase"
	@echo "  clean          - Clean up build and release directories"
