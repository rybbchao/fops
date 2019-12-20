package cmd

import (
	"fmt"

	"github.com/rybbchao/fops/pkg/fs"
	"github.com/spf13/cobra"
)

var mimeCmd = &cobra.Command{
	Use:   "mime",
	Short: "Print MIME type of file",
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath, _ := cmd.Flags().GetString("file")
		file, err := fs.ReadFile(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		mime, err := getMIMEType(filepath)
		if err != nil {
			return err
		}
		fmt.Println(mime)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mimeCmd)
	mimeCmd.Flags().StringP("file", "f", "", "Specify the file")
	mimeCmd.MarkFlagRequired("file")
}

func getMIMEType(filepath string) (string, error) {
	return fs.GetMIMEType(filepath)
}
