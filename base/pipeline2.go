package main

import (
	"fmt"
	"time"
)

func main() {
	n := make(chan int)
	s := make(chan int)
	go func() {
		for i := 0; i <= 20; i++ {
			n <- i
		}
		close(n)
	}()
	go func() {
		for i := range n {
			s <- i * i
			time.Sleep(250 * time.Millisecond)
		}
		close(s)
	}()
	for i := range s {
		fmt.Println(i)
	}
}
