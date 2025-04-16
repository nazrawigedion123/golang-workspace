package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addF(1.1, 2.2))
	fmt.Println(addT(1.1, 2))
}
