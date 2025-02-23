package main

import "fmt"

func main() {
	fmt.Println(foo(3))
	fmt.Println(bar(3, "three"))

}
func foo(x int) int {
	return x
}

func bar(x int, f string) (int, string) {
	return x, f
}
