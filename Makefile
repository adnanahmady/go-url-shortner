run:
	@go run main.go

test:
	@go test -v ./...
t: test

test.coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
tc: test.coverage

lint:
	@go vet ./...
	@gofmt -d -w .

build:
	@go build -o go-url-shortner main.go