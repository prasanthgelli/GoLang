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
	fmt.Println("I am", s.first, s.last, "- secretagent")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- secretagent")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I am called from Human,", h)
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
	p1 := person{
		first: "Dr.",
		last:  "Somebody",
	}
	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)
}
