package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, hashMap map[string]int) {
	println(file.Name())
	input := bufio.NewScanner(file)
	for input.Scan() {
		s := input.Text()
		fmt.Println(s)
		hashMap[s]++
	}
}

func main() {
	//创建map对象 返回的是一个指针
	hashMap := make(map[string]int)
	//读取文件参数
	files := os.Args[1:]
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "错误", err)
			continue
		}
		countLines(file, hashMap)
	}
	for line, n := range hashMap {
		fmt.Printf("%s的数量有%d个\n", line, n)
	}

}
