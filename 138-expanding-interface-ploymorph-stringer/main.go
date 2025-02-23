package main

import (
	"fmt"
	"log"
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

func logInfo(s fmt.Stringer) {
	log.Println("logged from 138", s.String())

}
func main() {
	b := book{
		title: "West with the light",
	}

	var c count = 42

	logInfo(b)
	logInfo(c)

}
