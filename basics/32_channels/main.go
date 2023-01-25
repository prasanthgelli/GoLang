package main

import "fmt"

func main() {
	c := make(chan int)
	// Every Channel Requires a go routine else it will be blocked
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	d := make(chan int, 3)
	go func() {
		d <- 42
		d <- 44
		d <- 45
	}()
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println("Directional Channels")
	// Can only read or can only pull
	e := make(chan<- int, 2)
	go func() {
		e <- 22
		e <- 24
	}()
	// <-e does'nt work as it is directional

}
