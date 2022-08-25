

setup:
	go install github.com/cosmtrek/air@latest

dev:
	air 

test:
	go install gotest.tools/gotestsum@latest
	gotestsum --format testname ./...
