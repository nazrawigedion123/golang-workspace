package main

import "fmt"

func main() {
	am := map[string]int{
		"buddy": 11,
		"air":   26,
	}
	fmt.Println(am)
	fmt.Println("the age of buddy was ", am["buddy"])

	delete(am, "buddy")
	delete(am, "ai")

	fmt.Println(am)

	// for range over a map
	for k, v := range am {
		fmt.Println(k, " - ", v)
	}

	for _, v := range am {
		fmt.Println(v)
	}

	for k := range am {
		fmt.Println(k)
	}

}
