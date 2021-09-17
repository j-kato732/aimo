FROM golang:1.17.0-buster

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /github.com/j-kato732/aimo

RUN apt update && apt install -y git curl protobuf-compiler tree vim \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/sqlite \
    && go get -u google.golang.org/grpc@latest \
# protoc-gen-go @ v1.27.1
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
# gepc @ v1.40.0
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    && go install github.com/infobloxopen/protoc-gen-gorm@latest \
    && go get github.com/grpc-ecosystem/grpc-gateway/v2


RUN git clone https://github.com/googleapis/googleapis.git ${GOPATH}/src \
    && git clone https://github.com/infobloxopen/protoc-gen-gorm.git ${GOPATH}/src/protoc-gen-gorm