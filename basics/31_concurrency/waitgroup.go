package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GOroutines\t", runtime.NumGoroutine())
	foo()
	bar()

}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}
