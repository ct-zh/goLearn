
GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc -I ./ --go_out=./ --micro_out=./ ./proto/user.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user cmd/main.go

.PHONY: test
test:
	go test -v ./.. -cover

.PHONY: dockerBuild
dockerBuild:
	docker build -t caot1995/user:latest .