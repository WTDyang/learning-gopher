package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hashMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		hashMap[input.Text()]++
	}
	for line, n := range hashMap {
		fmt.Printf("%s的数量有%d个\n", line, n)
	}
}
