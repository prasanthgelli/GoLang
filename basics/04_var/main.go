package main

import "fmt"

// Difference between var and := is
// := (short declaration operator) works only inside function body
// var works outiside the function body as well and if declared outside its scope is global

var abx = 99

// declare a variable of type int and assign a zero value
// false for booleans ,0 for integers, 0.0 for floats ,"" for strings
// nil for pointers, functions,  interfaces , slices, channels maps

var z int

func main() {
	// Short declaration operator
	// Declare a variable and assign a value(of certain type)
	x := 43
	fmt.Println(x)
	var y = 42
	fmt.Println(y)
	fmt.Println(abx)
}
