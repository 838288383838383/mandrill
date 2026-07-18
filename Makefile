VERSION := 0.1.0
LDFLAGS := -ldflags "-s -w"

.PHONY: build build-all release clean

build:
	go build $(LDFLAGS) -o mandrill ./cmd/mandrill/

build-all:
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/mandrill-linux-amd64 ./cmd/mandrill/
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/mandrill-linux-arm64 ./cmd/mandrill/
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/mandrill-darwin-amd64 ./cmd/mandrill/
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/mandrill-darwin-arm64 ./cmd/mandrill/
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/mandrill-windows-amd64.exe ./cmd/mandrill/
	@echo "Done! Binaries in dist/"

release: build-all
	gh release create v$(VERSION) \
		--title "v$(VERSION)" \
		--generate-notes \
		dist/mandrill-linux-amd64 \
		dist/mandrill-linux-arm64 \
		dist/mandrill-darwin-amd64 \
		dist/mandrill-darwin-arm64 \
		dist/mandrill-windows-amd64.exe

clean:
	rm -rf dist/ mandrill

install: build
	cp mandrill ~/.local/bin/ 2>/dev/null || mkdir -p ~/.local/bin && cp mandrill ~/.local/bin/
	@echo "Installed to ~/.local/bin/mandrill"

mac-sign:
	@echo "Signing darwin binaries with ad-hoc signature..."
	codesign --force --sign - dist/mandrill-darwin-amd64 2>/dev/null || true
	codesign --force --sign - dist/mandrill-darwin-arm64 2>/dev/null || true
	@echo "Done (ad-hoc signed)"
