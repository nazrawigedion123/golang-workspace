package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("if statement- x is %v , y is %v and ", x, y)

	//if statement
	if x < 4 && y < 4 {
		fmt.Println("both are less than 4")

	} else if x > 6 && y > 6 {
		fmt.Println("both are greater than 6")

	} else if x > 4 && y <= 6 {
		fmt.Println("x is greater than 4 and y is less than or equal to 6")

	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("non of the previous were met")
	}

	//switch statement

	fmt.Printf("switch statment- x is %v , y is %v and ", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("both are less than 4")
	case x > 6 && y > 6:
		fmt.Println("both are greater than 6")
	case x > 4 && y <= 6:
		fmt.Println("x is greater than 4 and y is less than or equal to 6")
	case y != 5:
		fmt.Println("y isnot 5")
	default:
		fmt.Println("non of the previous were met")
	}
}
