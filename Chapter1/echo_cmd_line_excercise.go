// Small program that prints the command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func main() {
	// excercise 1.1
	//fmt.Println(strings.Join(os.Args, " "))
	//excecise 1.2
	s, sep := "", " "
	for i, arg := range os.Args[1:] {
		s += strconv.Itoa(i) + sep + arg
		fmt.Println(s)
		s = ""
	}
}
