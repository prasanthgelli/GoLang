package main

import "fmt"

func main() {
	a := (4 == 2)
	b := (5 <= 10)
	c := (1 >= 9)
	d := (1 != 9)
	e := (1 > 9)
	f := (1 < 9)
	fmt.Println(a, b, c, d, e, f)
}
