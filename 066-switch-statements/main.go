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

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println(" this is default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is default case for x")
	}

	// no default fallthroug
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthroug : x is 41")
	case 42:
		fmt.Println("printed because of fallthroug : x is 42")
	case 43:
		fmt.Println("printed because of fallthroug : x is 43")
	default:
		fmt.Println("printed because of fallthroug :this is default case for x")

	}

	// no default fallthroug
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthroug : x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of fallthroug : x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of fallthroug : x is 43")
		fallthrough
	default:
		fmt.Println("printed because of fallthroug :this is default case for x")

	}

}
