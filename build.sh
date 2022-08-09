#!/bin/bash
rm -rf docs
swag init -g communication/http/http.go
go build

TARGET=dist/package/opt/posmaster/engine/bin
mkdir -p $TARGET
cp curly-engine $TARGET

