package main

func main() {
	ll := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	print("Sum of values is ,", foo(ll...))

}
func foo(x ...int) int {
	tot := 0
	for _, v := range x {
		tot += v
	}
	return tot

}
