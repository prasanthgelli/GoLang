// For this exercise
// 1. Create your own type. Have the underlying type be an int.
// 2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR”
// keyword
// 3. in func main
// 	a. print out the value of the variable “x”
// 	b. print out the type of the variable “x”
// 	c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
// 	d. print out the value of the variable “x”

package main

import "fmt"

type customtype int

var x customtype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
