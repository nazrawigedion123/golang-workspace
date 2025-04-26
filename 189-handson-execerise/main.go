package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	s := `[{
		"userId": 1,
		"id": 1,
		"title": "delectus aut autem",
		"completed": false
	  }]`
	bs := []byte(s)
	users := []User{}
	err := json.Unmarshal(bs,&users)
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println(users)
	}

}