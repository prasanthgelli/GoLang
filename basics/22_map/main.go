package main

import "fmt"

func main() {
	m := map[string]int{
		"james":   32,
		"maxpane": 22,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	// If key is not present it returns 0
	fmt.Println(m["adsa"])
	// Approach to find a key exists
	v, ok := m["adsa"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["james"]; ok {
		fmt.Println("This is the if print ", v)
	}
	// Add new element to map
	m["todd"] = 33
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
	// Delete a k,v in map
	delete(m, "todd")
	fmt.Println(m)

}
