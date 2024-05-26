package gofile

import (
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}
