package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))

}
func sum(li ...int) int {
	fmt.Println(li)
	fmt.Printf("%#v\n", li)
	sum := 0

	for _, l := range li {
		sum += l
	}

	return sum
}
