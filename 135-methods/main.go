package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("my name is ", p.first)

}

func main() {
	p1 := person{
		first: "blue",
	}
	p2 := person{
		first: "yellow",
	}
	p1.speak()
	p2.speak()
}
