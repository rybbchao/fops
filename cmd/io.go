package cmd

import (
	"fmt"
	"os"
)

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
