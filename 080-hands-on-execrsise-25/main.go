package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for y := 1; y <= 42; y++ {
		fmt.Printf("iteration %v-\t", y)
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Println("x is 0")
		case 1:
			fmt.Println("x is 1")
		case 2:
			fmt.Println("x is 2")
		case 3:
			fmt.Println("x is 3")
		default:
			fmt.Println("x is 4")

		}
	}
}
