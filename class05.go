package main 

import (
	"fmt"
)

var y = 42
var z = "sou uma variavel string"

func main() {
	fmt.Println(y) // ln = line
	fmt.Printf("%T\n", y) // f= format (formatcao)
	//%T = Type / Tipo
	fmt.Println("E agora?")
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

