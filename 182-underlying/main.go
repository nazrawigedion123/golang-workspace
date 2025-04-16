package main

import "fmt"

type myNumbers interface {
	~int | float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var n myAlias = 42

	fmt.Println(addT(n, 33))

}
