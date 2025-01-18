package main

import "fmt"

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

	if x < 42 && y < 42 {
		//&& requires to both true to run

		fmt.Println("both are less than the meaning of life")
	}
	if x > 30 || x < 42 {
		// || requires one of the statements to be correct
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

}
