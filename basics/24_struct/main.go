package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "james",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

	// Embedded Structs

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	fmt.Println(sa1, sa1.person.first, sa1.person.last, sa1.ltk)

	// anonymous struct
	prs1 := struct {
		first string
		last  string
	}{
		first: "james",
		last:  "bond",
	}
	fmt.Println(prs1)

}
