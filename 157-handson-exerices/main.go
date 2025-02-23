package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", substract)
	fmt.Printf("%T\n", doMath)

	fmt.Println(doMath(42, 16, add))
	fmt.Println(doMath(42, 16, substract))

}

func doMath(x int, y int, f func(int, int) int) int {
	return f(x, y)
}

func add(x int, y int) int {
	return x + y
}
func substract(x int, y int) int {
	return x - y
}
