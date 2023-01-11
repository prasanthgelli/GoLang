package main

import "fmt"

func main() {
	// Slice allows you to group values of same type
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	// Slice for range
	fmt.Println(len(x))
	//  Getting Elements at a particular index
	fmt.Println(x[2])
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slicing a slice

	fmt.Println(x[1:2])

	// Append to a slice

	x = append(x, 66, 77, 88, 122)
	fmt.Println(x)

	y := []int{96, 78, 52, 81}
	// ... is variadic parameter which means whole bunch of values in it
	x = append(x, y...)
	fmt.Println(x)

	// Deleting from a slice
	x = []int{1, 2, 3, 4, 5}
	x = append(x[:3], x[4:]...)
	fmt.Println(x)

	// Slice make
	a := make([]int, 10, 10) //(type,len,capacity)
	fmt.Println("\n\na")
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a[1] = 12
	//a[10] is not possible as the max len is 10 so we can use append
	a = append(a, 35)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// Multidimensional Slice
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println("\n\n\n", jb)
	mp := []string{"miss", "moneypenny", "strawberry", "hzlnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
