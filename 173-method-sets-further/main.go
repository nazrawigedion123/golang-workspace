package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is, ", d.first, " and am walking")
}

type yougin interface {
	walk()
	run()
}

func youngRun(y yougin) {
	y.run()
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is, ", d.first, " and am running")
}
func main() {

	d1 := dog{"henery"}
	d1.walk()
	d1.run()
	// youngRun(d1)

	d2 := &dog{"Padget"}
	youngRun(d2)

	d2.walk()
	d2.run()

}
