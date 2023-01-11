package main

import "fmt"

func main() {
	// for init; condition; post {
	// 	fmt.Println("Hello")
	// }
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("The outerloop %d and inner loop is %d\n", i, j)
		}

	}
}
