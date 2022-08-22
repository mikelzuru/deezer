# Common methodology based targets

.PHONY: prepare
prepare:

.PHONY: setup-dev
setup-dev:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.0

.PHONY: sanity-check
sanity-check:
	golangci-lint run ./...

.PHONY: build
build:

.PHONY: test
test:

.PHONY: release
release:

.PHONY: clean
clean:
