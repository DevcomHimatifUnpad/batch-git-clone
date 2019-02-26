#!/bin/bash
export GOBIN=$PWD/output 
echo $GOBIN
go install ./cmd/...