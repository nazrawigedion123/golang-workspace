package main

import (
	"fmt"
	"sort"
)

type People struct{
	Name string
	Age  int
}

func main() {

	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"james", "q", "m", "money penny", "doc. no"}

	fmt.Println("----------")
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("----------")
	fmt.Println(xs)
	sort.Strings(xs)

	fmt.Println(xs)

	people:=[]People{
		{ "bob",32},
		{"bonny",31},
		{"alex",41},
		{"blue",28},
	}

	fmt.Println("--------------------")
	fmt.Println(people)
}