package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 45
	b = 96
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// a = b we cannot do that as a and b are of different types
}
