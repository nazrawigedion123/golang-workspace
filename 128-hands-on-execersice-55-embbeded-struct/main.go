package main

import "fmt"

type engine struct {
	electric bool
}

type car struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	c1 := car{
		engine: engine{electric: true},
		make:   "byd",
		model:  "001",
		doors:  4,
		color:  "green",
	}

	c2 := car{
		engine: engine{electric: false},
		make:   "bmw",
		model:  "001",
		doors:  4,
		color:  "green",
	}

	fmt.Println(c1.engine, c1.make, c1.model, c1.doors, c1.color)

	fmt.Println(c2.engine, c2.make, c2.model, c2.doors, c2.color)

}
