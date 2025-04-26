package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "blue",
		Last:  "name",
		Age:   17}
	p2 := person{
		First: "Yellow",
		Last:  "name",
		Age:   17}
	people := []person{p1, p2}
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))

	}

}
