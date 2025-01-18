package main

import "fmt"

func main() {
	//for range loop
	// data structures -slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	//data structures -map
	m := map[string]int{
		"James":      42,
		"Momeypenny": 32,
	}

	for k, y := range m {
		fmt.Println("ranging over map", k, y)
	}

}
