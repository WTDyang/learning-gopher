package main

import (
	"fmt"
	"reflect"
)

type ReflectStruct struct {
	id   int    `ref:"Rid"`
	name string `ref:"Rname"`
}

func main() {
	ref := ReflectStruct{
		id:   1,
		name: "newOne",
	}
	v := reflect.ValueOf(ref)
	fmt.Printf("%T\n", v)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("ref")
		fmt.Println(name)
	}

}
