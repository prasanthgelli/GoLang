package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}

	if x := 42; x == 2 {
		fmt.Println("001")
	}
	fmt.Println("Other code ran")
	// fmt.Println(x) -> This does'nt work, but if x == 42 then it will work

	// IF else
	x := 42
	if x == 40 {
		fmt.Println("Our Value is 40")
	} else if x == 41 {
		fmt.Println("Our Value is 41")
	} else {
		fmt.Println("Our Value is not 40 or 41")
	}

	// Switch Statement
	a := 4
	switch {

	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println(("a is 3"))
	case a == 4:
		fmt.Println(("a is 4"))
	case a == 5:
		fmt.Println(("a is 5"))
	}

	// Expression Switches
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
