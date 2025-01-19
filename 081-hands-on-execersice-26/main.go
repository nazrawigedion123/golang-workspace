package main

import "fmt"

func main() {
	y := 0

	for y <= 20 {
		fmt.Println(y)

		y++
	}

	x := 0

	for {
		fmt.Println(x)
		if x >= 20 {
			break
		}
		x++
	}
}
