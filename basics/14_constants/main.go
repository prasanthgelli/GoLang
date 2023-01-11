package main

import "fmt"

const (
	a = 43
	b = 96.4454
	c = "James Bond"
)

// const (
// 	a int = 43
// 	b float64= 96.4454
// 	c string= "James Bond"
// )

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
