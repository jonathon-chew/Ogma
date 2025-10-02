package cmd

import (
	"fmt"
)

//This is the function to manage command line arguments
func Cmd(Arguments []string) {

	for _, v := range Arguments {
		fmt.Println("this is the call from the function, with argument:", v)
	}
}
