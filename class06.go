package main

import (
	"fmt"
)

var y = 42 
var z = `Sou uma variavel "string"` // Crase pode ser utilizado como """""
var a = `Berg disse
		Sou viado
         mas ele é meu "amigo"
`
var k = "Daniel Disse, que e foda. :-)"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	
	z = "43"
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(k)
	fmt.Printf("%T\n", k)
}

// RAW STRING
