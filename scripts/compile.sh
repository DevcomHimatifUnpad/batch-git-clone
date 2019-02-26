#!/bin/bash

echo "start compile to:"
GOBIN=$PWD/output go install ./cmd/...;
GOBIN=$PWD/output/ GOOS=windows GOARCH=amd64 go build ./cmd/...;
mv $PWD/*.exe $PWD/output;
echo $PWD/output; ls $PWD/output;
echo "finsihed compile";