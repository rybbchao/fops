package pkg

import (
	"fmt"
	"os"
)

func ReadFile(filepath string) (*os.File, error) {
	// Handle non-existent, invalid input file
	fi, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error: No such file '%s'", filepath)
		} else {
			return nil, err
		}
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("error: Expected file got directory '%s'", filepath)
	}
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
