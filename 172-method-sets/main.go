package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is, ", d.first, " and am walking")
}
func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is, ", d.first, " and am r unning")
}
func main() {

	d1 := dog{"henery"}
	d1.walk()
	d1.run()

	d2 := &dog{"Padget"}

	d2.walk()
	d2.run()

}
