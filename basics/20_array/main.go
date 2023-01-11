package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[2] = 22
	fmt.Println(x)
	// Length of array
	fmt.Println(len(x))

}
