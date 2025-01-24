package main

import "fmt"

func main() {
	// composit literal
	xs := []string{"a", "b", "c"}

	fmt.Printf("xs -  %#v \n", xs)

	//make a slice

	vi := make([]int, 0, 100)

	fmt.Printf("vi -  %#v  length - %v capacity - %v\n", vi, len(vi), cap(vi))

}
