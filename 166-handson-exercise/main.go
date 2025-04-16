package main

import (
	"fmt"
	"time"
)

func TimedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("elapsed time:", elapsed)
}
func MyFunction() {
	time.Sleep(2 * time.Second)
	fmt.Println("my function finished")
}
func main() {
	TimedFunction(MyFunction)
}
