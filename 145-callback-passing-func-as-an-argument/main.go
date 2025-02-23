package main

import "fmt"

func main() {
	a := 42
	b := 16
	fmt.Println("add", doMath(a, b, add))
	fmt.Println("substarct", doMath(a, b, substract))
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}
func substract(a int, b int) int {
	return a - b
}
