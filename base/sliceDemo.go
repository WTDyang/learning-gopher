package main

import "fmt"

func main() {
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = i
	}
	slice1 := slice[2:]
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	printSlice(slice)
	printSlice(slice2)
	slice2[2] = 100
	printSlice(slice)
	printSlice(slice2)
	reverse(slice)
	printSlice(slice)
}
func printSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf(" %d", slice[i])
	}
	fmt.Println("")
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
