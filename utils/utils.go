package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func ReadInput() (*os.File, error) {
	// Get the caller's file path
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("failed to get caller information")
	}

	// Construct the path to the input.txt file based on the caller's file location
	filePath := filepath.Join(filepath.Dir(filename), "input.txt") // Adjust the path as per your project's structure

	// Read the file
	return os.Open(filePath)
}
