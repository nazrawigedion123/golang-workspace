package main

import "fmt"

func main() {
	defer fmt.Println("after")
	defer fmt.Println("before")
	fmt.Println("before before")
}
