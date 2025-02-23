package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintf("the title is %v", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprintf(" the number is %v", strconv.Itoa(int(c)))
}
func main() {
	b := book{
		title: "West with the light",
	}

	var c count = 42

	fmt.Println(b)
	fmt.Println(c)

}
