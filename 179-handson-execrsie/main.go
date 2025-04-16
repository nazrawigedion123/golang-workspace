package main

import "fmt"

type Person struct {
	first string
}

func changeName(p Person) Person {
	p.first = "value"
	return p
}

func changeNamePointer(p *Person) {
	p.first = "pointer"
}

func main() {
	p1 := Person{first: "nazri"}
	fmt.Println(p1.first)
	p1 = changeName(p1)
	fmt.Println(p1)
	changeNamePointer(&p1)
	fmt.Println(p1)
}
