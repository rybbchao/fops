package fs

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func IsValidFile(filepath string) error {
	// Handle non-existent, invalid input file
	fi, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("No such file '%s'", filepath)
		} else {
			return err
		}
	}
	if fi.IsDir() {
		return fmt.Errorf("Expected file got directory '%s'", filepath)
	}
	return nil
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
