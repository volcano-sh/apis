update-codegen:
	./hack/update-codegen.sh

api-lint:
	go vet ./pkg/apis/...
	test -z "$$(gofmt -l pkg/apis/)"

build:
	go build ./...

test:
	go test ./...
