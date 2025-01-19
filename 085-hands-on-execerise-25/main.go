package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46}

	for i, y := range x {
		fmt.Println("range =", i, " value =", y)
	}
}
