package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	fmt.Printf("x is %v\t", x)
	// if statement
	if x <= 100 {
		fmt.Printf("and is bettween 0 and a 100")
	} else if x > 100 && x < 200 {
		fmt.Printf("and is between 100 and 200")
	} else {
		fmt.Printf("and is between 200 and 250")
	}
	for o := 0; o < 10; o++ {

		fmt.Println(rand.Intn(3))

	}

}
