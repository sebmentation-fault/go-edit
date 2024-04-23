package go_edit

type Model struct {
	GoEditFile []GoEditFile // we may have many files open at once
	CurrentFile int // index to array of which file we are looking at now
	GoEditMode GoEditMode // what mode is the user in?
	CurrentLine int // the line that thier cursor is on
	CurrentMaxWidth int // i.e. their cursor is at end of line, but line is only 4 char
	TermWidth int // the terminals width
	TermHeight int // the terminals height
	TopLine int // the topmost visible line
	BottomLine int // the bottommost visible line
}
