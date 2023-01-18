package main

import "fmt"

func main() {
	//1
	a := 95
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	//2
	x := 0
	fmt.Println("x befor", x)
	fmt.Println("x befor", &x)
	foo(&x)
	fmt.Println("x after", x)
	fmt.Println("x after", &x)
}
func foo(y *int) {
	fmt.Println("Y befor", y)
	fmt.Println("Y befor", *y)
	*y = 43
	fmt.Println("Y after", y)
	fmt.Println("Y after", *y)

}
