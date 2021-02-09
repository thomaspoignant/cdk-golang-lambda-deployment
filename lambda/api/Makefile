.PHONY: vendor
vendor:
	go mod vendor

test:
	go test -race ./...

build:
	go build -mod=vendor .

lambda-build:
	GOOS=linux go build -mod=vendor -o /asset-output/main
