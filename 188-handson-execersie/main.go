package main

import (
	"encoding/json"
	"fmt"
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

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

}
