package main

import (
	"fmt"
	"math"
)

type Square struct {
	width  float64
	length float64
}

func (s Square) area() float64 {
	return s.length * s.width
}

type Circle struct {
	radious float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radious, 2)
}

type Shape interface {
	area() float64
}

func info(s Shape) float64 {
	return float64(s.area())
}

func main() {
	s := Square{
		length: 3,
		width:  3,
	}
	c := Circle{
		radious: 3,
	}
	fmt.Println(info(s))
	fmt.Println(info(c))

}
