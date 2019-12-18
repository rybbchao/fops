package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version Info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fops " + version)
	},
}
