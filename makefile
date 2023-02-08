.PHONY: default
default: fmt test build

.PHONY: build
build:
	go build -o hx .

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: install
install:
	go install .
