.PHONY: proto
proto:
	protoc --go_out=./proto --micro_out=./proto ./proto/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hot cmd/main.go

.PHONY: test
test:
	go test -v cmd/main.go -cover

.PHONY: dockerbuild
dockerbuild:
	docker build -t caot1995/hot .
