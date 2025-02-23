package main

import "fmt"

func main() {
	foo()
	x := func() {
		fmt.Println("Anonymous func ran")
	}
	y := func(s string) {
		fmt.Println("this is me name", s)
	}
	x()
	y("Tood")

}
func foo() {
	fmt.Println("foo ran")
}
