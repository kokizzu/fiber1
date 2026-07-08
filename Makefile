
.PHONY: setup dev test verify-dependency-security


setup:
	go install github.com/cosmtrek/air@latest

dev:
	air 

test:
	go install gotest.tools/gotestsum@latest
	gotestsum --format testname ./...

verify-dependency-security:
	bash ./scripts/verify-dependency-security.sh
