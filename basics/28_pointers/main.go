package main

import "fmt"

func main() {
	a := 95
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// var b int := a --> This is not possible as they are two diff types one is int and other is *int
	var b *int = &a
	// Reference to an address
	fmt.Println(b)
	// * derefencing the pointer
	fmt.Println(*b)
	fmt.Println(&a)  //-> This gives address
	fmt.Println(*&a) // this gives the value

	*b = 44
	fmt.Println(a)

	// Pass Values using addresses
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
