package main

import "fmt"

func main() {
	xi := []int{42, 43}

	xi = append(xi, 44, 45)
	fmt.Println(xi)

}
