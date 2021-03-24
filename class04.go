package main

import (
	"fmt"
)

var y int = 43

var z int
var w string
var a bool
var b float64

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	foo()
	fmt.Println("Encerrando Foo...")
	foo()
}

func foo() {
	fmt.Println("novamente", y)
	fmt.Println("o valor de z é:", z)
	fmt.Println("o valor de w é:", w)
	fmt.Println("o valor de a é:", a)
	fmt.Println("o valor de b é:", b)
}