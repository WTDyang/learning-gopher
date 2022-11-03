package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 3
	T := reflect.TypeOf(i)
	fmt.Println(T.String())
	fmt.Printf("%#v\n%[1]T", T)
	fmt.Println("\n------------------------")
	V := reflect.ValueOf(i)
	fmt.Println(V.String())
	fmt.Printf("%#v\n%[1]T", V)
}
