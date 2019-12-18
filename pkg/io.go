package pkg

import (
	"fmt"
	"net/http"
	"os"
)

func ReadFile(filepath string) (*os.File, error) {
	// Handle non-existent, invalid input file
	fi, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("No such file '%s'", filepath)
		} else {
			return nil, err
		}
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("Expected file got directory '%s'", filepath)
	}
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetMIMEType(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}
