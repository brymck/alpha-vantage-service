OWNER := brymck

SERVICE_NAME := $(notdir $(CURDIR))
GO_FILES := $(shell find . -name '*.go')
PROTO_PATH := /usr/local/include
PROTO_FILES := $(shell find proto -name '*.proto')
GENPROTO_FILES := $(patsubst proto/%.proto,genproto/%.pb.go,$(PROTO_FILES))

all: proto test build

init: .init.stamp

.init.stamp:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go mod download
	touch $@

proto: $(GENPROTO_FILES)

genproto/%.pb.go: proto/%.proto | .init.stamp
	mkdir -p $(dir $@)
	protoc -Iproto -I$(PROTO_PATH) --go_out=plugins=grpc:genproto $<

test: profile.out

profile.out: $(GO_FILES) $(GENPROTO_FILES) | .init.stamp
	go test -race -coverprofile=profile.out -covermode=atomic ./...

build: service

client: $(GO_FILES) $(GENPROTO_FILES) | .init.stamp
	go build -ldflags='-w -s' -o client cmd/client/*.go

service: $(GO_FILES) $(GENPROTO_FILES) | .init.stamp
	go build -ldflags='-w -s' -o service cmd/web/*.go

run: service
	./service

docker:
	docker build . --tag docker.pkg.github.com/$(OWNER)/$(SERVICE_NAME)/$(SERVICE_NAME)

clean:
	rm -rf genproto/ .init.stamp profile.out client service

.PHONY: all init proto test build run docker clean
