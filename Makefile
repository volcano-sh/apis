update-codegen:
	./hack/update-codegen.sh

build:
	go build ./...

test:
	go test ./...