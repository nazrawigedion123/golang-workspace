package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}
func spliceDelta(li *[]int) {
	*li = []int{3, 4, 5}
}
func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3}
	fmt.Println(xi)
	spliceDelta(&xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["james"] = 32
	fmt.Println(m)
	mapDelta(m, "james")
	fmt.Println(m)
}
