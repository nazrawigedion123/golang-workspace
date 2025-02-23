package main

import "fmt"

func main() {

	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
}

func sum(i []int) int {
	total := 0
	for _, o := range i {
		total += o
	}
	return total
}
