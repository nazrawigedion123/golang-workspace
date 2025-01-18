package main

import "fmt"

func main() {
	//sequence
	x := 40
	y := 5
	fmt.Println("x=%v \n y=%v \n", x, y)

	//loops / iterations
	//for statements
	/*
		The go for loop is similar to but not the same as C's.
		It unifies for and while and there is no do while.
		There are three forms, only one fo which had semicolons.

		Like a C for
		for inint; condition; post{}

		like a C while
		for condition {}

		//likr a C for {;;}

		for {}


	*/
	// https://go.dev/doc/effective_go#for
	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to five: %v \t first for loop \n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}
	// break
	// takes you out of the loop

	for {
		fmt.Printf("y is %v \t\t\t third loop \n", y)
		if y > 20 {
			break
		}
		y++

	}

	// continue
	//takes to next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

}
