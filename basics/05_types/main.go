package main

import "fmt"

var y = 42
var z string = "Hey!!!"

// Whatever is inside “ comes directly even the ""
var a string = `He said,"Hey Man Wassup!"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	//z = 43
	//z = 43, this cannot be allowed as it was assigned to string already an integer is not accepted
	// GO is a static programming language
	// Not a dynamic programming languag

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
