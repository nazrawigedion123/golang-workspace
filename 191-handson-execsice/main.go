package main

import (
	"fmt"
	"sort"
)

type People struct {
	name string
	age  int
}

// ByName implements sort.Interface based on the Name field.
type ByName []People

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].name < a[j].name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByAge implements sort.Interface based on the Age field.
type ByAge []People

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	people := []People{
		{"fares", 8},
		{"bob", 18},
	}
	xi:=[]int{3,2}
	fmt.Println(people,xi)

	sort.Ints(xi)

	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)

	



}