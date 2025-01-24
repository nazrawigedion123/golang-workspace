package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	for i, v := range a {
		fmt.Printf("%v - %v - %T\n", i, v, v)
	}

}
