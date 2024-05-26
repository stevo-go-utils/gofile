package gofile

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{
		path: path,
	}
}

func (f File) Path() string {
	return f.path
}

func (f File) Name() string {
	return f.path
}
