package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(s string) {
		fmt.Println("this is me name", s)
	}("blue")

}
func foo() {
	fmt.Println("foo ran")
}
