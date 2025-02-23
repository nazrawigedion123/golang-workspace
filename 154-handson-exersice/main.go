package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() string {
	return "my name is " + p.first
}

func main() {
	x1 := person{
		first: "james",
		age:   24,
	}
	fmt.Println(x1.speak())
}
