package main

import "fmt"

//new 只会进行初始化 并讲内存空间清零 这样的化切片内部的大小和容量都是0了，当然不能添加元素了
//make 除了分配内存之外 还进行了初始化
func main() {
	newDemo()
}
func newDemo() {
	Slice := new([]int)

	(*Slice)[0] = 1
	fmt.Printf("%d", (*Slice)[0]) //panic: runtime error: index out of range [1] with length 0

	//*Slice = append(*Slice, 1)
	//fmt.Printf("%d", (*Slice)[0]) //1
}
func makeDemo() {
	Slice := make([]int, 1, 10)
	Slice[0] = 1
	fmt.Printf("%d", Slice[0])
}
