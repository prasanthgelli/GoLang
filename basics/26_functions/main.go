package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Data")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
	// Variadic parameter
	val := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(val)

	// Anonymous Func
	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("Anonymous func ran with value x->", x)
	}(42)

}

// func( r reveiver) identifier(parameters) (returns(s)){...}
func foo() {
	fmt.Println("This is from foo")
}
func bar(s string) {
	fmt.Println("Hello", s)

}
func woo(st string) string {
	return fmt.Sprint("Hello from woo,", st)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `,says "hello"`)
	b := true
	return a, b

}

// Variadic parameter
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
