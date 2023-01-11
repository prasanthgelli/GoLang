package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
	for i := 33; i < 122; i++ {
		fmt.Printf("%v\t%x\t%#U\n", i, i, i)
	}
}
