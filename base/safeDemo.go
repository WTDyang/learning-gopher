package main

import (
	"fmt"
	"sync"
)

var (
	mu  sync.Mutex
	num int
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		index := i
		wg.Add(1)
		go func(index int) {
			mu.Lock()
			defer mu.Unlock()
			num++
			wg.Done()
		}(index)
	}
	wg.Wait()
	fmt.Println(num)
}
