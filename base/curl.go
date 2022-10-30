package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "访问失败", err)
			os.Exit(500)
		}
		body, err1 := ioutil.ReadAll(resp.Body)
		code := resp.StatusCode
		if err1 != nil {
			fmt.Fprintln(os.Stderr, "读取失败", err)
		}
		resp.Body.Close()
		fmt.Printf("状态码：%d\n%s\n", code, body)
	}
}
