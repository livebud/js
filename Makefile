GO_SOURCE := ./...

default: test

ci: fmt vet test

fmt:
	@ test -z "$(shell go fmt $(GO_SOURCE))"

vet:
	@ go vet $(GO_SOURCE)

test:
	@ go test --race --failfast --timeout=20m $(GO_SOURCE)
