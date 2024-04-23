package go_edit

import (
	"fmt"
	"log"
	"os"

	xterm "golang.org/x/term"
)

type TermModel struct {
	fdesciptor int
	state xterm.State
	width, height int
}

// Initialises a `TermModel` struct.
// Retruns nil if:
//	- the user is not in a terminal
func InitUI() *TermModel {
	fd := int(os.Stdout.Fd())
	if !xterm.IsTerminal(fd) {
		log.Print("Called outside terminal setting")
		return nil
	}

	state, err := xterm.MakeRaw(fd)

	if err != nil {
		log.Print("The terminal emulator could not be taken over")
		return nil
	}

	width, height, err := xterm.GetSize(fd)

	if err != nil {
		log.Print("Could not get size of terminal emulator")
	}

	return &TermModel{
		fdesciptor: fd,
		state: *state,
		width: width,
		height: height,
	}
}

// Given a `TermModel`, a `GoEditFile` and a `GoEditMode`, refresh the terminal
// terminal window, drawing on all the text, and updating the infobar as appropriate
func (termModel *TermModel) DrawWindow(goEditFile *GoEditFile, status int) {
	infoWidth := termModel.width - 2

	// HEADER
	fmt.Print("+")
	for i := 1; i < infoWidth; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	// INFORBAR
	fmt.Print("| ")
	fname := "FILE NAME"
	switch status {
	case EDITING_MODE:
		fmt.Print(fname)
		for i:= 2 + len(fname); i < infoWidth - 8; i++ {
			fmt.Print(" ")
		}
		fmt.Print("| [ESC] ")
		break
	case ESCAPED_MODE:
		fmt.Print(fname)
		for i:= 2 + len(fname); i < infoWidth - 18; i++ {
			fmt.Print(" ")
		}
		fmt.Print("| [Q]uit | [S]ave ")
		break
	}
	fmt.Print("|")

	fmt.Print("+")
	for i := 1; i < infoWidth; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	// MAIN
	for i := 2; i < termModel.height - 1; i++ {
		fmt.Print("| ")
		for j := 0; j < infoWidth; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	// FOOTER
	fmt.Print("+")
	for i := 1; i < infoWidth; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
