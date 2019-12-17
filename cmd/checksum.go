package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"

	"github.com/spf13/cobra"
)

// add flag name and hash method here
var mapHashFunc = map[string]func() hash.Hash{
	"md5":    md5.New,
	"sha1":   sha1.New,
	"sha256": sha256.New,
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		file, err := readFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// sort the output
		sortedHashName := []string{"md5", "sha1", "sha256"}
		for _, hashName := range sortedHashName {
			// check if the flag is true
			if check, err := cmd.Flags().GetBool(hashName); err == nil {
				if check {
					val, err := checksum(file, mapHashFunc[hashName]())
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(hashName + ":\t" + val)
				}
			} else {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringP("file", "f", "", "Specify the file")
	// add the flag dynamically
	for hashName, _ := range mapHashFunc {
		checksumCmd.Flags().Bool(hashName, false, fmt.Sprintf("Print %s checksum", hashName))
	}
	checksumCmd.MarkFlagRequired("file")
}

func checksum(reader io.Reader, h hash.Hash) (string, error) {
	if _, err := io.Copy(h, reader); err != nil {
		return "", err
	}
	// convert to hex string
	return hex.EncodeToString(h.Sum(nil)), nil
}
