package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "martini", "chocolate"}
	jm := []string{"jenny", "moneypenny", "guiness", "wolverine tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
}
