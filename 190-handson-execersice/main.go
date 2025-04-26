package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	users := []User{
		{1, "bob", 19},
		{1, "blue", 20},
	}
	fmt.Println(users)
	

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("something wrong",err)
	} 
}
