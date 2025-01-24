package main

import "fmt"

func main() {
	am := map[string]int{
		"buddy": 11,
		"air":   26,
	}
	fmt.Println(am)
	fmt.Println("the age of buddy was ", am["buddy"])

	// delete(am, "buddy")
	// delete(am, "ai")

	fmt.Println(am)

	if y, ok := am["buddy"]; ok {
		fmt.Println("buddy: ", y)
	} else {
		fmt.Println("buddy doesnt exist")
	}

}
