# fops

`fops` is a simple command-line application which is built using [Cobra](https://github.com/spf13/cobra)

## Introduction

### Structure

This project is based on [Cobra](https://github.com/spf13/cobra).

```
cmd/
    root.go
    <othersCommand>.go
main.go
```

The `main.go` initilizes the Cobra. And each command has its own file in `cmd/`. For example, The `version` command is written in `cmd/version.go`.

### checksum

All the algorithms supported by `checksum` and theirs implementations are stored in a map. Therefore, it's easy to extend the flags. 

## Features
- **linecount**: Implementation of `wc -l [file]`
- **checksum**: Print the checksum of the file. (*--md5*, *--sha1*, *--sha256*)

## Run

```
go run main.go
```

## Test

```
go test ./...
```

## Build

```
go build -o fops
```

## TODO

- `linecount` needs to check if the file is a binary file
