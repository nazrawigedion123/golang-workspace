package main

import "fmt"

func foo(s string) {
	fmt.Println(s)
}

func main() {
	foo("fooo dear")
	bar("nazrawi")
	fmt.Println(aloha("nazrawi"))
	fmt.Println(dogYears("buddy", 7))

}

func bar(s string) {
	fmt.Println("my name is ", s)
}

func aloha(s string) string {
	return fmt.Sprint("they call me ", s)
}
func dogYears(name string, age int) (string, int) {
	return fmt.Sprint(name, " is "), age * 7
}
