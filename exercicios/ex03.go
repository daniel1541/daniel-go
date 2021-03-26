// 1. Utilizando o operador para declaração curta (ou short) de variável,
// ATRIBUA os seguintes VALORES para os IDENTIFICADORES "x", "y" and "z":

// a. 42
// b. "James Bond"
// c. true

// 2. Agora imprima os valores armazenados nestas variáveis utilizando:
// a. Uma única chamada de Print (Um único Print).
// b. Múltiplas chamadas de Print.

package main

import (
	"fmt"
)

var a = 42
var b = "James Bond"
var c = true

func main() {

	fmt.Println("O valor de A", a, "O valor de B", b, "O Valor de C", c)
	// No exercício você pode ver que na letra a) nós pedimos uma única chamada de Print..
	// e isso você fez aí acima...
	// Mas faltou a letra b. Múltiplas chamadas de Print.
	// Seria algo mais ou menos assim:
	fmt.Println("O valor de a é", a)
	fmt.Println("O valor de b é", b)
	fmt.Println("O valor de c é", c)
}
