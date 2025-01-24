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
	an["george"] = 78

	fmt.Printf("%#v\n", an)

	// for range over a map
	for k, v := range an {
		fmt.Println(k, " - ", v)
	}

	for _, v := range an {
		fmt.Println(v)
	}

	for k := range an {
		fmt.Println(k)
	}

	//for range over a slice

	xi := []int{42, 43, 44}

	for k, v := range xi {
		fmt.Println(k, " - ", v)
	}
	for _, v := range xi {
		fmt.Println(" - ", v)
	}
	for k := range xi {
		fmt.Println(k, " - ")
	}

}
