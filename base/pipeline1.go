package main

import (
	"fmt"
	"time"
)

func main() {
	n := make(chan int)
	s := make(chan int)
	go counter(n)
	go squarer(n, s)
	for {
		if i, ok := <-s; ok {
			fmt.Println(i)
		} else {
			break
		}
	}

}
func counter(n chan int) {
	for i := 0; i <= 20; i++ {
		time.Sleep(250 * time.Millisecond)
		n <- i
	}
	close(n)
	fmt.Println("over")
}
func squarer(n chan int, s chan int) {
	for {
		if i, ok := <-n; ok {
			s <- i * i
		} else {
			fmt.Println("结束了!")
			break
		}
	}
	close(s)
}
