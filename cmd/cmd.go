package cmd

import (
	"fmt"
	"os"
)

//This is the function to manage command line arguments
func Cmd(Arguments []string) {

	for _, v := range Arguments {
		if v == "--version" || v == "-v"{
			versionNumber := "0.0.1"
			fmt.Printf("Version %s", versionNumber)
			os.Exit(0)
		} else {
			fmt.Println("this is the call from the function, with argument:", v)
		}
	}
}
