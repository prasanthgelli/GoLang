package main

import "fmt"

const (
	_ = iota
	a = iota + 2023
	b = iota + 2023
	c = iota + 2023
	d = iota + 2023
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
