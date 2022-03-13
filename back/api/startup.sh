#!/usr/bin/env bash

nohup go run grpc/server/main.go >> log/server.log & 
go run grpc/gateway/main.go >> log/gateway.log
