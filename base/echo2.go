package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	sep = " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}
