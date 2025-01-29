package main

import "fmt"

func main() {
	x1 := []int{1, 2, 3, 4, 5, 6}
	x := sum(x1...)
	fmt.Println("sum = ", x)

}
func sum(i ...int) int {
	n := 0
	for _, v := range i {
		n += v
	}
	return n
}
