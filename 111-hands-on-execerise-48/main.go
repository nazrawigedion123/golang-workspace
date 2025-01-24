package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "Shake, not stired"}
	jm := []string{"jenny", "moneypenny", "am 008"}

	ss := [][]string{jb, jm}
	for i, v := range ss {
		for ii, vv := range v {
			fmt.Println(i, "-", v, " ", ii, " ", vv)
		}
	}

}
