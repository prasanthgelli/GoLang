package main

import "fmt"

func foo() {
	fmt.Println("This is foo")
}
func bar() {
	fmt.Println("This is bar")
}

func main() {
	defer foo()
	bar()
}
