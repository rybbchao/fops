# fops

`fops` is a simple command-line application which is built using [Cobra](https://github.com/spf13/cobra).

## Introduction

- Based on [Cobra](https://github.com/spf13/cobra).
- Use [gabriel-vasile/mimetype](https://github.com/gabriel-vasile/mimetype) to detect the file type.
- All the algorithms supported by `checksum` and theirs implementations are stored in a map. Therefore, it's easy to extend the flags.
- Use build time flags to set the application version.

### Structure

This project is based on [Cobra](https://github.com/spf13/cobra).

```
cmd/
    root.go
    <commands>.go
pkg/
    <utils>
tests/
    <input files for test>
main.go
```

The `main.go` initilizes the Cobra. And each command has its own file in `cmd/`. For example, The `version` command is written in `cmd/version.go`.

## Features
- **linecount**: Implementation of `wc -l [file]`.
- **checksum**: Print the checksum of the file. (*--md5*, *--sha1*, *--sha256*).

## Run

```
go run main.go
```

## Test

```
go test ./...
```

## Build

```bash
chmod +x build.sh
./build.sh $version
```
