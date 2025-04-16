package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "drop by drop, the bucket gets filled"
	q := "persistently, patiently , you are bound to succeed "
	r := "the meaning of life is"
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("value \t\ttype\tvalue\n")
	fmt.Printf("%v\t%T\t%v\n", a, a, *a)
	fmt.Printf("%v\t%T\t%v\n", b, b, *b)
	fmt.Printf("%v\t%T\t%v\n", c, c, *c)
	fmt.Printf("%v\t%T\t%v\n", d, d, *d)

}
