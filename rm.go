package gofile

import "os"

func Rm(path string) (err error) {
	return os.Remove(path)
}

func RmRf(path string) (err error) {
	return os.RemoveAll(path)
}
