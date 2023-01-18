package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a, b, c)
}
func foo() int {
	return 99
}
func bar() (int, string) {
	return 89, "This is bar"

}
