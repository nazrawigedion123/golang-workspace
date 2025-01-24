package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type agent struct {
	person
	ltk bool
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Moneypenny",
		last:  "Jenny",
		age:   28,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Println(p1.last, " ", p1.first)

	j1 := agent{
		person: p1,
		ltk:    true,
	}
	fmt.Println(j1)
	fmt.Println(j1.person)

}
