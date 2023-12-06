package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadInputAsFile() (*os.File, error) {
	// Get the caller's file path
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("failed to get caller information")
	}

	// Construct the path to the input.txt file based on the caller's file location
	filePath := filepath.Join(filepath.Dir(filename), "input.txt")

	// Read the file
	return os.Open(filePath)
}

func ReadInputAsLineSlice() ([]string, error) {
	// Get the caller's file path
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("failed to get caller information")
	}

	// Construct the path to the input.txt file based on the caller's file location
	filePath := filepath.Join(filepath.Dir(filename), "input.txt")

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func SplitByLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}
