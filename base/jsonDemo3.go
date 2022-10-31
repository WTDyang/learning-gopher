package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	student1 := student{
		Id:   0,
		Name: "韩墨霖",
		Sid:  "20201202048",
	}
	fmt.Printf("%#v\n", student1)
	//json.NewEncoder()
	sJson, err := json.MarshalIndent(student1, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", sJson)

}
