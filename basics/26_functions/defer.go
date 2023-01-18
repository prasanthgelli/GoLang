package main

import "fmt"

// defer is used to temporarly skip the function calling
// it will be called at the end of the execution
// Use case open a file and immediately defer close the file
// This will make sure that once the execution of the scope is completed we close the file

func main() {
	defer foo()
	bar()
}
func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
