package main

import "fmt"

type person struct {
	first string
}
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("my name is ", p.first)

}
func (p secretAgent) speak() {
	fmt.Println("im secret agent ", p.first)

}

type human interface {
	speak()
}

func saySomeThing(h ...human) {
	for _, v := range h {

		v.speak()
	}

}
func main() {
	p1 := person{
		first: "blue",
	}
	p2 := secretAgent{
		person: p1,
		ltk:    true,
	}
	p1.speak()
	p2.speak()
	saySomeThing(p1, p2)
}
