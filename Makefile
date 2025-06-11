.PHONY: run lint test commit

run:
	@echo "Starting server with Air hot reload..."
	@air

lint:
	golangci-lint run

test:
	go test ./...

commit:
	cz c