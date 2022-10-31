package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Sid  string `json:"studentId,omitempty"`
}

func main() {
	student1 := []Student{
		Student{
			Id:   0,
			Name: "韩墨霖",
			Sid:  "",
		},
		Student{
			Id:   0,
			Name: "马钰鹏",
			Sid:  "20201202075"},
	}
	fmt.Printf("%#v\n", student1)
	sJson, err := json.MarshalIndent(student1, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", sJson)
	var name []struct {
		Name string `json:"name"`
	}
	json.Unmarshal(sJson, &name)
	fmt.Printf("%v \n", name)
}
