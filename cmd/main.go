package main

import (
	"fmt"
	"log"
	"os"
go_edit "sebmentation-fault/go-edit/pkg"
)

func main() {
	fmt.Print("Enter the filepath\n> ")
	var fPath string
	_, err := fmt.Scanln(&fPath)
	fmt.Println()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	_, statusCode := go_edit.ReadFile(fPath)

	if statusCode == nil {
		log.Println("Expected a status code from go_edit.ReadFile")
		os.Exit(-2)
	}

	switch statusCode.Code {
	case go_edit.FILE_DOES_EXIST:
		fmt.Print("File does not exist. Would you like to create it? [Y/n] ")
		var res rune
		_, err := fmt.Scanf("%c", &res)
		fmt.Println()

		if err != nil {
			fmt.Println(err)
			os.Exit(-3)
		}

		switch res {
		case 'y':
		case 'Y':
		case '\n': // create the file
			osFile, err := os.Create(fPath)
			osFile.Close()

			if err != nil {
				fmt.Println(err)
				os.Exit(-4)
			}

			// create a GoEditFile struct
			gef := go_edit.GoEditFile{
				Path: fPath,
				Buffer: "",
				IsModified: false,
			}

			term := go_edit.InitUI()

			term.DrawWindow(&gef, go_edit.EDITING_MODE)

			break

		case 'n':
		case 'N':
			fmt.Println("Exiting without opening file.")
			os.Exit(0)
			break
		}

		break

	case go_edit.SUCCESS:
		fmt.Println("Opening file...")
		gef := go_edit.GoEditFile{
			Path: fPath,
			Buffer: "",
			IsModified: false,
		}

		term := go_edit.InitUI()

		term.DrawWindow(&gef, go_edit.EDITING_MODE)

		break
	}

	os.Exit(0)
}
