package main

import "fmt"

func main() {
	//1. Print every number from 1 to 10,000
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}

	// 2. print all uppercase alphabets 3 times each
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	// 3. Years I'm alive using code block ---for condition {}
	year := 1998
	for year < 2024 {
		fmt.Println(year)
		year++
	}

	// 4. Years I'm alive using code block for {}
	year = 1998
	for {
		if year > 2023 {
			break
		}
		fmt.Println(year)
		year++
	}

	// 5. remainder between 10 to 100 divided by 4
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

	// 8. Switch with no conditions
	switch {
	case true:
		fmt.Println("This should be printed")
	case false:
		fmt.Println("This should not be printed")
	}

	// 9. Switch with variable

	favsport := "golf"
	switch favsport {
	case "cricket":
		fmt.Println("Favsport is cricket")
	case "football":
		fmt.Println("Favsport is football")
	case "golf":
		fmt.Println("Favsport is golf")
	}

	// 10. Conditions
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
