#!/bin/bash

go build -ldflags "-X main.Version=$1" -o fops