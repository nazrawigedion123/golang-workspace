package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old. \n", name, age)

	fmt.Printf("%s is %d years old \t and the type is %T and %T. \n", name, age, name, age)
	fmt.Println(`
				❤️❤️❤️❤️😭😭😭
				miss u forever
				loveley blablabla`)

	var g int
	fmt.Println(g)
	g = 42
	fmt.Println(g)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%v, %b ,%#X \n", a, a, a)
	fmt.Printf("%v, %b ,%#X \n", b, b, b)
	fmt.Printf("%v, %b ,%#X \n", c, c, c)
	fmt.Printf("%v, %b ,%#X \n", e, e, e)
	fmt.Printf("%v, %b ,%#X \n", d, d, d)
	fmt.Printf("%v, %b ,%#X \n", f, f, f)

}
