package main

import "fmt"

type PointA struct {
	X, Y int
}
type PointB struct {
	X, Y int
}
type Point struct {
	PointA
	PointB
}

func (p PointA) dis() int {
	return p.X - p.Y
}
func (p PointB) dis() int {
	return p.X - p.Y
}

func main() {
	point := Point{
		PointA: PointA{
			X: 1,
			Y: 2,
		},
		PointB: PointB{
			X: 4,
			Y: 8,
		},
	}
	//fmt.Println(point.dis())	//不明确的引用
	fmt.Println(point.PointB.dis())
}
