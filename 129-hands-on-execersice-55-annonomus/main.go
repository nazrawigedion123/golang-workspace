package main

import "fmt"

func main() {

	p1 := struct {
		first   string
		friends map[string]int
		favD    string
	}{
		first:   "blue",
		friends: map[string]int{"yello": 12, "red": 90},
		favD:    "martini",
	}
	fmt.Println(p1)

}
