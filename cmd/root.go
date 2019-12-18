package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version string = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:   "fops",
	Short: "File Ops",
}

func SetVersion(v string) {
	version = v
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
