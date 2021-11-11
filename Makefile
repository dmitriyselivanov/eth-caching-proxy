.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

test:
	go test ./... | { grep -v 'no test files'; true; }

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf eth-caching-proxy
	go clean -i .

docker:
	docker-compose up -d --build