package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		for y := 5; y > 0; y-- {
			fmt.Printf("%v of ", i)
			fmt.Println(y)
		}
	}
}
