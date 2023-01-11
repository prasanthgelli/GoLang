package main

import "fmt"

var y string
var z int

func main() {
	// Declare a variable to be of certain type
	// and assign a value of that type to the variable
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Adding Data Here"
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println("Printing the value of z", z, "ending")
	fmt.Printf("%T\n", z)
}
