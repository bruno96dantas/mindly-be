build:
	@go build -o api cmd/main.go

run:
	@echo "Starting server..."
	@air

lint:
	@golangci-lint run

test:
	go test ./...

commit:
	cz c