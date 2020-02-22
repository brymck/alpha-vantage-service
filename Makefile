PROTOS := example

SERVICE_NAME := $(notdir $(CURDIR))
GO_FILES := $(shell find . -name '*.go')
PROTO_PATH := /usr/local/include
PROTO_DIR := proto

all: dep proto test build

dep: .dep.stamp

.dep.stamp: go.mod go.sum
	go get -u github.com/golang/protobuf/protoc-gen-go
	go mod download
	touch $@

proto: genproto/.dirstamp

genproto/.dirstamp: dep
	mkdir -p genproto
	protoc -Iproto -I$(PROTO_PATH) --go_out=plugins=grpc:genproto proto/*.proto
	touch $@

test: profile.out

profile.out: proto $(GO_FILES)
	go mod download
	go test -race -coverprofile=profile.out -covermode=atomic ./...

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
