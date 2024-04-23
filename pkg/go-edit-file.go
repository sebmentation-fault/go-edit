package go_edit

type GoEditFile struct {
	Path string
	Buffer string
	IsModified bool
}

type GoEditMode int // Either EDITING_MODE or ESCAPED_MODE

const (
	// Means that the user is editing the file, can type to write to the buffer, etc
	EDITING_MODE GoEditMode = iota
	// Means that the user is in escaped mode, can perform actions such as quiting, saving, etc.
	ESCAPED_MODE
)

// Get a convenient string for the file name
func (file *GoEditFile) ToString() string {
	if file == nil {
		return "No File"
	}

	return file.Path
}

