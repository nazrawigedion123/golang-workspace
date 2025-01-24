package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

type foo int

func main() {

	var a foo = 12
	fmt.Println("a =", a)

	p1 := person{
		first:    "James",
		last:     "Bond",
		icecream: []string{"chocolate", "chocolate chip"},
	}

	p2 := person{
		first:    "Moneypenny",
		last:     "Jenny",
		icecream: []string{"vinilla"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Println(p1.last, " ", p1.first)

	for _, v := range p1.icecream {
		fmt.Println("fav icecream flavour for ", p1.first, " - ", v)
	}
	for _, v := range p2.icecream {
		fmt.Println("fav icecream flavour for ", p2.first, " - ", v)
	}

}
