.PHONY:  build fmt
build:
	go build .

fmt:
	go fmt -w .

test:
	go test -cover .

vet:
	go vet .
