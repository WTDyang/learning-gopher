package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Ref ReflectStruct

func (r Ref) get() {
	fmt.Println("方法get")
}
func (r Ref) set() {
	fmt.Println("方法set")
}

func main() {
	x := Ref{
		id:   1,
		name: "newOne",
	}
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}
