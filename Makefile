PROTOS := example

SERVICE_NAME := $(notdir $(CURDIR))
GO_FILES := $(shell find . -name '*.go')
PROTO_DIR := proto

all: proto test build

proto: genproto/.dirstamp

genproto/.dirstamp:
	mkdir -p genproto
	go mod download
	protoc --go_out=plugins=grpc:genproto -I proto proto/*.proto
	touch $@

test: coverage.txt

coverage.txt: $(GO_FILES)
	go mod download
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

build: service

service: proto $(GO_FILES)
	go mod download
	go build -o service .

run: service
	./service

docker:
	docker build . --tag $(SERVICE_NAME)

clean:
	rm -f service coverage.txt

.PHONY: all proto test build run docker clean
