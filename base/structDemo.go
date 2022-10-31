package main

import "fmt"

type circle struct {
	x int
	y int
	r int
}
type wheel struct {
	circle
	nr int
}

func main() {
	w := wheel{
		circle: circle{
			x: 1,
			y: 2,
			r: 0,
		},
		nr: 3,
	}
	fmt.Println(w.x)
	fmt.Println(w.circle.y)
	fmt.Printf("%#v \n", w)
}
