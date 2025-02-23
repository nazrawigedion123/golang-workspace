package main

import "fmt"

func main() {
	result := x()            // Call x, get the inner function
	finalResult := result()  // Call the inner function, get 42
	fmt.Println(finalResult) // Print 42
	//or in one line
	fmt.Println(x()())
}

func x() func() int {
	return func() int {
		return 42
	}
}
