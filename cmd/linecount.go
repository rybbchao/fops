package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		file, err := readFile(filename)
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

func readFile(filename string) (*os.File, error) {
	// Handle non-existent, invalid input file
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error: No such file '%s'", filename)
		} else {
			return nil, err
		}
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("error: Expected file got directory '%s'", filename)
	}
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
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
