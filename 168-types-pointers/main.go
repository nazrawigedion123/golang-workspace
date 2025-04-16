package main

import "fmt"

func main() {
	x := 2
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%v\t%T\n", &x, &x)
}
