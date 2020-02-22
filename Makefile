PROTOS := example

SERVICE_NAME := $(notdir $(CURDIR))
GO_FILES := $(shell find . -name '*.go')
PROTO_DIR := proto

all: proto test build

proto:
	mkdir -p genproto
	protoc --go_out=plugins=grpc:genproto -I $(PROTO_DIR) $(PROTO_DIR)/*.proto

test: coverage.txt

coverage.txt: $(GO_FILES)
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

build: service

service: $(GO_FILES)
	go mod download
	go build -o service .

run: service
	./service

docker:
	docker build . --tag $(SERVICE_NAME)

clean:
	rm -f service coverage.txt

.PHONY: all proto test build run docker clean
