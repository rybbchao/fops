package cmd

import (
	"bytes"
	"fmt"
	"io"

	"github.com/rybbchao/fops/pkg"
	"github.com/spf13/cobra"
)

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("file")
		file, err := pkg.ReadFile(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		count, err := checkLineCount(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(count)
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
