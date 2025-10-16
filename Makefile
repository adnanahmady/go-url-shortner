run:
	@go run main.go

test:
	@go test -v ./...
t: test

lint:
	@go vet ./...
	@gofmt -d -w .

build:
	@go build -o go-url-shortner main.go