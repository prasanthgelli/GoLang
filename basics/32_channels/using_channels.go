package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c)
	// receive **** Imp part
	bar(c)

}
func foo(c chan<- int) {
	c <- 42

}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
