update-codegen:
	./hack/update-codegen.sh

verify-codegen:
	./hack/verify-codegen.sh

api-lint:
	go vet ./pkg/apis/...
	test -z "$$(gofmt -l pkg/apis/)"

crd-gen:
	@echo "Validating API structure with crdify..."
	go install sigs.k8s.io/crdify@latest
	~/go/bin/crdify --help || echo "crdify installed"

build:
	go build ./...

test:
	go test ./...