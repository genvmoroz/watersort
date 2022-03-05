include linting.mk

OSNAME=$(shell go env GOOS)

.PHONY: deps
deps:
	go mod tidy
	go mod download
	go mod verify

.PHONY: unit_test
unit_test:
	go test -v -cover `go list ./...` -count=1

.PHONY: bench
test:
	TODO

# The Linux build uses UPX to reduce the final binary size by ~70%
.PHONY: build
build:
	GOOS=$(OSNAME) \
		go build \
		-ldflags="-s -w" \
		-o artifacts/svc \
		.
	@if [[ "$(OSNAME)" == "linux" ]]; then \
		mv ./artifacts/svc ./artifacts/svc-unpacked && \
		upx -q -o ./artifacts/svc ./artifacts/svc-unpacked && \
		rm ./artifacts/svc-unpacked ; \
	fi
