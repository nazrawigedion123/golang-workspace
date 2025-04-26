package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password"
	bs,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println(bs)
	}

	login:= "password"
	err= bcrypt.CompareHashAndPassword(bs,[]byte(login))
	if err!=nil{
		fmt.Println("ur pass word is wrong")
	}else{
		fmt.Println("correct")
	}
}