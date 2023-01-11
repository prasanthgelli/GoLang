package main

import "fmt"

const (
	a         = 42
	b float64 = 96
)

func main() {

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
