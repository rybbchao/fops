package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/rybbchao/fops/pkg/fs"
	"github.com/spf13/cobra"
)

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath, _ := cmd.Flags().GetString("file")
		err := fs.IsValidFile(filepath)
		if err != nil {
			return err
		}
		isBinary, err := fs.IsBinary(filepath)
		if err != nil {
			return err
		}
		if isBinary {
			return fmt.Errorf("Cannot do linecount for binary file '%s'", filepath)
		}
		file, err := os.Open(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		count, err := checkLineCount(file)
		if err != nil {
			return err
		}
		fmt.Println(count)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringP("file", "f", "", "Specify the file")
	linecountCmd.MarkFlagRequired("file")
}

func checkLineCount(reader io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := reader.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return 0, err
		}
	}
}
