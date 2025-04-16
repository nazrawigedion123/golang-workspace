package main

import "fmt"

func main() {
	x := 2
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(x)
}
