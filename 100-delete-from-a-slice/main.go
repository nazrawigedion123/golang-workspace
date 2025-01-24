package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("xi -  %#v \n", xi)

	//delete from slice

	xi = append(xi[:3], xi[4:]...)
	fmt.Printf("deleted 4 xi -  %#v \n", xi)

}
