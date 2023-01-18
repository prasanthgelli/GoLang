package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretagent struct {
	person
	ltk bool
}

func (s secretagent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretagent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	sa2 := secretagent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

}
