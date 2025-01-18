package main

import (
	"fmt"

	"math/rand"
)

func main() {
	//sequence
	fmt.Println("first statement to run")
	fmt.Println("second statement to run")
	x := 40 //third statement to run
	y := 5  // fourth statement to run
	fmt.Printf("x=%v \n y=%v \n", x, y)

	//conditional
	//if statements
	//switch statements

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater than or equal to %v \n", z, x)
	} else {
		fmt.Printf(" z id %v and that is less than x whis is %v \n", z, x)
	}

}
