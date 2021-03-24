package main

import (
	"fmt"
)

var y string = "Bond, James Bond"
var z int

func main()  {
	fmt.Println("Imprimindo o valor de y", y, "e pronto.")
	fmt.Printf("%T\n", y)
	fmt.Println("Imprimindo o valor de z", z, "e pronto.")
	fmt.Printf("%T\n", z)
}