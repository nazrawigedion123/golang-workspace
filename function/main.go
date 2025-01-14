package main

import "fmt"

func add(x, y int) int {

	return (x + y)
}
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println(add(42, 13))
	sayHello()
	fmt.Println(swap("hello", "world"))

}
func sayHello() {
	fmt.Println("hello")
}
