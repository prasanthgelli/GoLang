package main

import "fmt"

// - for 1.
type person struct {
	first          string
	last           string
	fav_ice_creame []string
}

// - for 3.
type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxary bool
}

func main() {
	// 1.
	fmt.Println("----------------------------------------")
	p1 := person{
		first:          "James",
		last:           "Bond",
		fav_ice_creame: []string{"choco", "Vanialla"},
	}
	p2 := person{
		first:          "max",
		last:           "payne",
		fav_ice_creame: []string{"butterscotch", "hzlnut"},
	}
	fmt.Println(p1, p2)
	for i, v := range p1.fav_ice_creame {
		fmt.Println(i, v)
	}
	for i, v := range p2.fav_ice_creame {
		fmt.Println(i, v)
	}
	// 2.
	fmt.Println("----------------------------------------")

	m := map[string]person{}
	m[p1.last] = p1
	m[p2.last] = p2
	fmt.Println(m)
	for i, v := range m[p1.last].fav_ice_creame {
		fmt.Println(i, v)
	}
	for i, v := range m[p2.last].fav_ice_creame {
		fmt.Println(i, v)
	}

	// 3.
	fmt.Println("----------------------------------------")
	v1 := vehicle{
		doors: 4,
		color: "black",
	}
	fmt.Println(v1.doors, v1.color)
	t1 := truck{
		vehicle: vehicle{
			doors: 6,
			color: "white",
		},
		fourwheel: true,
	}
	fmt.Println(t1.vehicle.doors, t1.vehicle.color, t1.fourwheel)
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxary: true,
	}
	fmt.Println(s1.vehicle.doors, s1.vehicle.color, s1.luxary)

	// 4.
	fmt.Println("----------------------------------------")

	anomStruct := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   44,
	}
	fmt.Println(anomStruct)
}
