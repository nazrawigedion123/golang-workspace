package main

import "fmt"

/* variatic function that take a slice of int*/

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(sum2([]int{1, 2, 3, 4, 5, 6, 7, 8}))

}
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total

}

func sum2(x []int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total

}
