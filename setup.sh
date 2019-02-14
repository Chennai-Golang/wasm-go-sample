#!/usr/bin/env bash

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
GOOS=js GOARCH=wasm go build -o lib.wasm main.go
go build server.go
