package main

import (
	"github.com/rybbchao/fops/cmd"
)

var Version string = "v0.0.1"

func main() {
	cmd.SetVersion(Version)
	cmd.Execute()
}
