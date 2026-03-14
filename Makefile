VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"
DIST    := dist
TMPL    := packaging/mcpb/manifest.json.tmpl

.PHONY: build test cover lint vet clean mcpb mcpb-all

build:
	go build $(LDFLAGS) -o biblium ./cmd/biblium

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run ./...

vet:
	go vet ./...

# Build a single .mcpb bundle.
#   make mcpb GOOS=darwin GOARCH=arm64
#   make mcpb GOOS=windows GOARCH=amd64
mcpb:
	$(eval BINARY := $(if $(filter windows,$(GOOS)),biblium.exe,biblium))
	$(eval MCPB_PLATFORM := $(if $(filter windows,$(GOOS)),win32,$(GOOS)))
	@mkdir -p $(DIST)/mcpb-staging/server
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(DIST)/mcpb-staging/server/$(BINARY) ./cmd/biblium
	@sed -e 's|{{VERSION}}|$(VERSION)|g' \
	     -e 's|{{BINARY}}|$(BINARY)|g' \
	     -e 's|{{PLATFORM}}|$(MCPB_PLATFORM)|g' \
	     $(TMPL) > $(DIST)/mcpb-staging/manifest.json
	@cd $(DIST)/mcpb-staging && zip -qr ../biblium-$(GOOS)-$(GOARCH).mcpb .
	@rm -rf $(DIST)/mcpb-staging
	@echo "Built: $(DIST)/biblium-$(GOOS)-$(GOARCH).mcpb"

# Build .mcpb bundles for all supported platforms.
mcpb-all:
	@$(MAKE) mcpb GOOS=darwin  GOARCH=arm64
	@$(MAKE) mcpb GOOS=windows GOARCH=amd64
	@$(MAKE) mcpb GOOS=linux   GOARCH=amd64
	@$(MAKE) mcpb GOOS=linux   GOARCH=arm64

clean:
	rm -f biblium biblium.exe coverage.out
	rm -rf $(DIST)/
