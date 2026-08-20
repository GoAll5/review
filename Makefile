BINARY_NAME=app

.PHONY: build run test clean

build:
	go build -o bin/$(BINARY_NAME) cmd/review/main.go

run: build
	./bin/$(BINARY_NAME)

test:
	go test -v ./...

clean:
	rm -rf bin/