package fs

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
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

func IsBinary(filepath string) (bool, error) {
	mime, err := mimetype.DetectFile(filepath)
	if err != nil {
		return false, err
	}
	return mime.Is("application/x-mach-binary"), nil
}

func GetMIMEType(filepath string) (string, error) {
	mime, err := mimetype.DetectFile(filepath)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}
