package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
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

}
