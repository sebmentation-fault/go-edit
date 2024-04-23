package go_edit

type GoEditFile struct {
	Path string
	Buffer string
	IsModified bool
}

func (gef *GoEditFile) ToString() string {
	if gef == nil {
		return "No File"
	}

	return gef.Path
}
