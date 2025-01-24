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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v)
		for _, v1 := range v.icecream {
			fmt.Println(v.first, v.last, v1)
		}

	}

}
