package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Flags struct {
	IgnoreFolders []string
}

// This is the function to manage command line arguments
func Cmd(Arguments []string) Flags {

	//(#2) TODO: Add in the ability to exclude certain files

	//(#4) TODO: Add in the ability to ignore errors

	var FlagArguments Flags

	for numberOfArguments := 0; numberOfArguments < len(Arguments); numberOfArguments++ {
		flag := Arguments[numberOfArguments]

		switch flag {

		case "--version", "-v":
			versionNumber := "v0.0.4"
			fmt.Printf("Version %s", versionNumber)
			os.Exit(0)

		case "--ignore", "-i":
			if numberOfArguments+1 >= len(Arguments) {
				log.Print("[ERROR]: no file found after -i flag")
				return FlagArguments
			}

			for i := numberOfArguments + 1; i < len(Arguments); i++ {
				if !strings.HasPrefix(Arguments[i], "-") {
					FlagArguments.IgnoreFolders = append(
						FlagArguments.IgnoreFolders,
						Arguments[i],
					)
					numberOfArguments++
				} else {
					break
				}
			}

		default:
			log.Println("Unable to deal with argument:", flag)
		}
	}

	return FlagArguments
}
