package main

import (
	"fmt"
	"math"
)

func main() {
	x := powanitor(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powanitor(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
