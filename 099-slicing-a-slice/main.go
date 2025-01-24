package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("xi -  %#v \n", xi)

	//[inclucive: exclusive]

	fmt.Printf("[0: 4] xi -  %#v \n", xi[0:4])

	//[:execlusive]
	fmt.Printf("[:7] xi - %#v \n", xi[:7])

	//[inclusive:]
	fmt.Printf("[4:] xi - %#v\n", xi[4:])

	//[:]
	fmt.Printf("[:] xi - %#v\n", xi[:])

}
