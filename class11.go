package main

import (
	"fmt"
)

var a int 

type cerveja int

var b cerveja

func main() {
	a = 42
	fmt.Println("Olá imprimindo o valor de A:", a)
	fmt.Printf("%T\n", a)

	b = 43
	fmt.Println("Olá imprimindo o valor de B:", b)
	fmt.Printf("%T\n", b)

	a = b 
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
