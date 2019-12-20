#!/bin/bash

go build -ldflags "-X github.com/rybbchao/fops/cmd.Version=$1" -o fops