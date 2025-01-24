package main

import "fmt"

func main() {
	xs := map[string][]string{"bond_james": {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dir":           {"cats", "ice cream", "sunsets"}}
	fmt.Printf("%#v\n", xs)
	xs["flaming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range xs {
		for k1, v1 := range v {
			fmt.Printf("%v -%v - %v\n", k, k1, v1)
		}
	}
	delete(xs, "flaming_ian")
	fmt.Println("---------------------record deleted------------------------------")

	for k, v := range xs {
		for k1, v1 := range v {
			fmt.Printf("%v -%v - %v\n", k, k1, v1)
		}
	}

}
