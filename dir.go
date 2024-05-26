package gofile

import (
	"errors"
	"io/fs"
	"os"
)

func MkDir(path string) error {
	if Exists(path) {
		return errors.New("dir already exists")
	}
	return os.Mkdir(path, os.ModePerm)
}

func ForceMkDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func MkDirAndParents(path string) error {
	if Exists(path) {
		return errors.New("dir already exists")
	}
	return os.MkdirAll(path, os.ModePerm)
}

func ForceMkDirAndParents(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

type ReadDirOpts struct {
	suffix string
	filter ReadDirFilterType
}

type ReadDirOptFunc func(*ReadDirOpts)

type ReadDirFilterType string

const (
	ReadDirDirFilter  ReadDirFilterType = "dir"
	ReadDirFileFilter ReadDirFilterType = "file"
	ReadDirNoneFilter ReadDirFilterType = "none"
)

func DefaultReadDirOpts() ReadDirOpts {
	return ReadDirOpts{
		suffix: "",
		filter: ReadDirNoneFilter,
	}
}

func filterFilesSuffix(files []fs.DirEntry, suffix string) (resFiles []fs.DirEntry) {
	for _, file := range files {
		fn := file.Name()
		if len(fn) < len(suffix) {
			continue
		}
		if fn[len(fn)-len(suffix):] == suffix {
			resFiles = append(resFiles, file)
		}
	}
	return resFiles
}

func filterFilesType(files []fs.DirEntry, fileType ReadDirFilterType) (resFiles []fs.DirEntry) {
	if fileType == ReadDirNoneFilter {
		return files
	}
	if fileType == ReadDirFileFilter {
		for _, file := range files {
			if !file.IsDir() {
				resFiles = append(resFiles, file)
			}
		}
	} else if fileType == ReadDirDirFilter {
		for _, file := range files {
			if file.IsDir() {
				resFiles = append(resFiles, file)
			}
		}
	}
	return resFiles
}

func ReadDir(path string, opts ...ReadDirOptFunc) (files []fs.DirEntry, err error) {
	rdOpts := DefaultReadDirOpts()
	for _, opt := range opts {
		opt(&rdOpts)
	}

	unfilteredFiles, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}

	if len(rdOpts.suffix) != 0 {
		unfilteredFiles = filterFilesSuffix(unfilteredFiles, rdOpts.suffix)
	}
	files = filterFilesType(unfilteredFiles, rdOpts.filter)

	return files, nil
}

func ReadDirNames(path string, opts ...ReadDirOptFunc) (fileNames []string, err error) {
	rdOpts := DefaultReadDirOpts()
	for _, opt := range opts {
		opt(&rdOpts)
	}

	unfilteredFiles, err := os.ReadDir(path)
	if err != nil {
		return fileNames, err
	}

	var files []fs.DirEntry
	if len(rdOpts.suffix) != 0 {
		unfilteredFiles = filterFilesSuffix(unfilteredFiles, rdOpts.suffix)
	}
	files = filterFilesType(unfilteredFiles, rdOpts.filter)

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
