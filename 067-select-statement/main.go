package main

import (
	"fmt"

	"math/rand"
	"time"
)

func main() {

	// conditional
	// if statements
	//switch statements

	//concurrency
	//select a channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get int64, convert it in to type time.Duration, the use it in time.Sleep
	// Int63n returns and int 64
	//type duration's underlying type is int64
	// time.Sleep takes type duration
	//time.Millisecond id a constant in the time package
	//https://pkg.go.dev/timepkg-constants
	d1 := time.Duration(rand.Int63n(250))

	d2 := time.Duration(rand.Int63n(250))
	//fmt.Printf("%v \t %T\n",d1,d2)
	//time.Sleep(d1 * time.Millisecond)
	//fmt.Println("done")
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

}
