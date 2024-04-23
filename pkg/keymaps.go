package go_edit

import go_edit "sebmentation-fault/go-edit/pkg"

type Action int

const (
	NOT_SUPPORTED Action = -1

	// actions when in escaped mode
	QUIT Action = iota
	SAVE
	ENTER_EDITING

	// actions when in editing mode
	ENTER_ESCAPED
	PUT_CHAR_AT_CURSOR
	CREATE_NEWLINE
	DELETE_CHAR
	MOVE_CURSOR_LEFT
	MOVE_CURSOR_RIGHT
	MOVE_CURSOR_DOWN
	MOVE_CURSOR_UP
)

const (
	keyEsc rune = 27
	keyDel rune = "\x1b[3~"
	keyUp rune = "\x1b[A"
	keyDown rune = "\x1b[B"
	keyLeft rune = "\x1b[C"
	keyRight rune = "\x1b[D"
)

func GetAction(mode go_edit.GoEditMode, key rune) Action {

	switch mode {
	case ESCAPED_MODE:
		switch key {
		case 'q':
		case 'Q':
			return QUIT
		case 's':
		case 'S':
			return SAVE
		case '\r':
		case '\n':
			return ENTER_EDITING
		}
	case EDITING_MODE:
		switch key {
		case keyEsc:
			return ENTER_ESCAPED
		case '\r':
		case '\n':
			return CREATE_NEWLINE
		case keyDel:
			return DELETE_CHAR
		}
		return PUT_CHAR_AT_CURSOR
	}

	return NOT_SUPPORTED
}

