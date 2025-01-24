package main

import "fmt"

func main() {
	am := map[string]int{
		"buddy":   11,
		"nazrawi": 26,
	}
	fmt.Println("the age of buddy was ", am["buddy"])

	fmt.Println(am)

	fmt.Printf("%#v \n", am)

	an := make(map[string]int)
	an["blue"] = 32
	an["green"] = 56

	fmt.Printf("%#v\n", an)

}
