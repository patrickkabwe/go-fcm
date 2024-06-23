GO ?= go

.PHONY: test
test:
	@$(GO) test -v -cover -covermode=atomic -coverprofile  coverage.out ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

test-coverage:
	@$(GO) tool cover -html=coverage.out

clean:
	@$(GO) clean -x -i ./...