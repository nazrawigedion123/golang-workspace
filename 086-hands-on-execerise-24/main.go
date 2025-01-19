package main

import "fmt"

func main() {

	m := map[string]int{
		"james":      42,
		"Moneypenny": 43,
	}

	for i, y := range m {
		fmt.Printf("name %v age %v \n", i, y)
	}

	if age, ok := m["james"]; !ok {
		fmt.Println("ages for james is ", age)
	}
	if age, ok := m["Q"]; !ok {
		fmt.Println("age for Q is ", age)
	} else {
		fmt.Println("age for Q doesnt exist ", age)

	}
}
