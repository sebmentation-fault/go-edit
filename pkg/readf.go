package go_edit

import (
	"os"
)

type ReadCode struct {
	Code int
}

var SUCCESS = 0
var FILE_DOES_EXIST = 1

// Function to read a file and if successfull, then parse
// the contents of the file to a TedFile struct
func ReadFile(path string) (*GoEditFile, *ReadCode) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		code := ReadCode{Code: FILE_DOES_EXIST,}
		return nil, &code
	}

	goEditFile := GoEditFile{
		Path: path,
		Buffer: string(bytes),
		IsModified: false,
	}

	return &goEditFile, &ReadCode{Code: SUCCESS}
}

