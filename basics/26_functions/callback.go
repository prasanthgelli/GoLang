package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ii)
	s1 := sum(ii...)
	fmt.Println("Sum of all numbers is -> ", s1)
	s2 := even_sum(sum, ii...)
	fmt.Println("Sum of even numbers is ->", s2)

}
func sum(data ...int) int {
	total := 0
	for _, v := range data {
		total += v
	}
	return total
}
func even_sum(f func(xx ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}
