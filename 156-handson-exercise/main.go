package main

import "fmt"

func main() {
	total := Add(3, 4)

	fmt.Println(total)
}

func Add(x int, y int) int {
	return (x + y)

}
