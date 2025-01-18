package main

import (
	"fmt"
)

func main() {
	fmt.Println("first statment to run")
	fmt.Println("second statment to run")
	x := 40
	y := 5
	fmt.Printf("x=%v\n y=%v\n", x, y)
	//conditional
	//if statments
	//switch statements

	if x < 42 {
		fmt.Println("less than the meaning of lif")
	}
	if x < 42 {
		fmt.Println("less than the meaning of life")
	} else {
		fmt.Println("greater than or equal to the meaning of life")
	}

	if x < 42 {
		fmt.Println("less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to meaning of life")

	} else {
		fmt.Println("greater than to the meaning of life")
	}
}
