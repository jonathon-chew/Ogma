package cmd

import (
	"fmt"
)

func Cmd(Arguments []string) {
	//This is the function

	for _, v := range Arguments {
		fmt.Println("this is the call from the function, with argument:", v)
	}
}
