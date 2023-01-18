package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   32,
	}
	p2 := person{
		First: "miss",
		Last:  "moenypenny",
		Age:   30,
	}
	people := []person{p1, p2}
	fmt.Println(p1)
	fmt.Println(p2)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
