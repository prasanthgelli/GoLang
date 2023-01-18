package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("First Func Expression")
	}
	f()

	fint := func(y int) {
		fmt.Println("First Func Expression with params y->", y)
	}
	fint(45)

	// Return using func
	s1 := foo()
	fmt.Println(s1)

	dat := bar()
	fmt.Printf("%T\n", dat)
	val := dat()
	fmt.Println(val)
}

func foo() string {
	s := "hello worlds"
	return s
}

func bar() func() int {
	return func() int {
		return 1000
	}
}
