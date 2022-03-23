#! /bin/bash

# Build go as shared library
go build -buildmode=c-shared -o libgoex.so felis.est/examples/cpp_go
