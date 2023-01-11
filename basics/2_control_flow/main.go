package main

import "fmt"

func main() {
	fmt.Println("Hello I'm Learning Control Flow")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	fmt.Println("This is foo")
}
func bar() {
	fmt.Println("Exited hurray!!!!")
}

// Types of control flow
// 1. Sequential
// 2. Conditional
// 3. Iterative
