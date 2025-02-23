package main

import (
	"bytes"
	"fmt"
)

// type person struct{
// 	first string

// }

//	func (p person) writeOut(w io.Writer) error {
//		_, err := w.Write([]byte(p.first))
//		return err
//	}
func main() {

	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()
	// s := []byte("hello gophers!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	b := bytes.NewBufferString("hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("its firday here")
	fmt.Println(b.String())
	b.Write([]byte("happy happy"))
	fmt.Println(b.String())
}
