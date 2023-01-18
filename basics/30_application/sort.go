package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{7, 4, 9, 2, 0, 2, 5, 7, 1, 1, 5}
	xs := []string{"abc", "def", "book", "sample", "data", "final"}

	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println("Sorted Slice are")
	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)

}
