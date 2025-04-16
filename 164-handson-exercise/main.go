package main

import "fmt"

func main() {

	x := 2
	y := 2

	fmt.Println(domath(x, y, add))
	fmt.Println(domath(x, y, substract))
}

func domath(x int, y int, f func(int, int) int) int {
	return f(x, y)

}
func add(x int, y int) int {
	return x + y
}
func substract(x int, y int) int {
	return x - y
}
