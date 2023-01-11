package main

import "fmt"

func main() {
	// 1. Create array of length 5 and assign value and iterate using range and print the values, print the type of array
	var arr [5]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)

	// 2.Create slice of length 10 and assign value and iterate using range and print the values, print the type of array
	fmt.Println("----------------------------")
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slc {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slc)

	// 3.
	fmt.Println("----------------------------")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)

	// 4.
	fmt.Println("----------------------------")
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	//5.
	fmt.Println("----------------------------")
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = append(x[:3], x[6:]...)
	fmt.Println(y)

	// 6.
	fmt.Println("----------------------------")
	mk := make([]string, 50, 50)
	ab := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	for i, v := range ab {
		mk[i] = v
	}
	fmt.Println(mk)
	fmt.Println(len(mk))
	fmt.Println(cap(mk))

	// 7.
	fmt.Println("----------------------------")
	sl1 := []string{"James", "Bond", "Shaken, not stirred"}
	sl2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	m := [][]string{sl1, sl2}
	for i, v := range m {
		fmt.Println(i)
		for j, val := range v {
			fmt.Print("\t", j, "\t", val, "\n")
		}
	}

	//8.
	fmt.Println("----------------------------")
	mapdata := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(mapdata)
	// 9.
	fmt.Println("----------------------------")
	mapdata["elonmusk"] = []string{`twitter`, `spacex`, `mars`}
	for k, v := range mapdata {
		fmt.Println(k, v)
	}
	//10.
	fmt.Println("----------------------------")
	delete(mapdata, "elonmusk")
	fmt.Println(mapdata)
}
