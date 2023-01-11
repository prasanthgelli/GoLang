package main

import "fmt"

var y = 42

func main() {
	//printf is formatting pring
	//println is printing in new line
	//print returns int and err code
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#X\n", y)

	//Sprintf store the print statement String to a variable
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	//%v is default value printing
	fmt.Printf("%v", y)
}
