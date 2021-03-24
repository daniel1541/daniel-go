package main 

import (
	"fmt"
)

var y = 42

func main()  {
	fmt.Println("Ola Mundo")
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Printf("%b\n", y) //binario
	fmt.Printf("%x\n", y) //hexadecimal
	fmt.Printf("%#x\n", y) //hexadecimal com 0x
	fmt.Printf("%X\n", y) //hexadecimal em MAIUSCULO

	x := "Berg queima"
	fmt.Printf("%#x\n%b\n%x\n", x, x, x)

}

