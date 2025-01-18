package main

import "fmt"

func main() {
	x := 83 / 40
	y := 83 % 40

	fmt.Println(x, y)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("This is an even number ", i)
		}
	}

}
