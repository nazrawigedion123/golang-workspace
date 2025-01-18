package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("from the init")

}
func main() {
	x := rand.Intn(250)

	fmt.Printf("x is %v\t", x)

	// switch
	switch {
	case x <= 100:
		fmt.Printf("and is bettween 0 and a 100")
	case x > 100 && x < 200:
		fmt.Printf("and is between 100 and 200")
	default:
		fmt.Printf("and is between 200 and 250")

	}

}
