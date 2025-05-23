package fs

import (
	"fmt"
	"os"
)

func ReadFileString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Unable to read file %s", filename)
	}
	return string(data), nil
}

func WriteFileString(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("Unable to write file %s", filename)
	}
	return nil
}

