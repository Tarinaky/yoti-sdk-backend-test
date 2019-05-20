container: test
	docker build -t ysbt_api -f Dockerfile ..

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ysbt_api .

test:
	go test ./...

.phony: container build test
