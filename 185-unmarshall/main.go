package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	s := `[{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}]`
	bs := []byte(s)

	people := []People{}
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Printf("%T \n%T\n", s, bs)
	fmt.Println("all the data ",people)
	for i,v:= range people{
		fmt.Println(i," - ", v.ID, " - ", v.Title)

	}
	

	
}
