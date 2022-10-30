package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("开始。。。")
	for i := range os.Args {
		fmt.Println(os.Args[i])
	}
}
