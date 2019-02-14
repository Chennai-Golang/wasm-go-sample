#!/usr/bin/env bash

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

# TODO: Compile main.c to main.wasm

GOOS=js GOARCH=wasm go build -o lib.wasm main.go
go build server.go
