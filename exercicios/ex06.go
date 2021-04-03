// 1. Crie um programa que, no escopo de pacote crie as seguintes variáveis:
// - nome (string)
// - idade (int)
// - peso (float64)
// - pesopizza (int)
// 
// 	a. para nome, atribua seu nome
// 	b. para idade, atribua sua idade
// 	c. para peso, atribua seu peso
// 
// 2. Na função main:
// 
// Imprima o seguinte parágrafo da forma que você desejar,
// respeitando as seguintes regras:
// 
// 	a. Onde você encontrar <nome>,
// 	deverá ser impresso o valor da variável nome. Onde encontrar <idade>, deverá
// 	ser impresso o valor da variável idade. E assim por diante.
// 	b. Onde você encontrar <tipo_variável_nome>, deverá ser impresso o tipo da variável
// 	nome e assim por diante.
// 
// 
// Meu nome é <nome>. Estou aprendendo o básico de Go e git.
// Possuo <idade> anos e estou pesando aproximadamente <peso>.
// Utilizei variáveis para lhe passar estas informações pessoais com Go.
// Por exemplo, utilizei a variável nome, com o tipo <tipo_variável_nome>.
// Utilizei outras duas variáveis, para idade e peso, com os tipos <tipo_variável_idade> e <tipo_variável_peso>.
// Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.
// 
// 3. Após imprimir o texto acima, atribua o valor da variável pesopizza de forma que ele seja o seu peso + 1kg.
// Em seguida imprima o valor e o tipo da variável pesopizza


package main

import (
	"fmt"
)

var nome string = "daniel"
var idade int = 27
var peso float64 = 80
var pesopizza float64 = 1 

func main() {

	var total = peso + pesopizza

	fmt.Println("Meu nome é", nome ,". Estou aprendendo o básico de Go e git.")
	fmt.Println("Possuo", idade , "anos e estou pesando aproximadamente", peso , "Kg.")
	fmt.Println("Utilizei variáveis para lhe passar estas informações pessoais com Go.")
	fmt.Printf("Por exemplo, utilizei a variável nome, com o tipo %T\n", nome)
	fmt.Printf("Utilizei outras duas variáveis, para idade e peso, com os tipos %T e %T\n", idade, peso)
	fmt.Println("Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.")
	fmt.Println("Em seguida imprima o valor com a pizza", total)
	fmt.Printf("e o tipo da variável é %T\n", total)
	
}	

// Não conseguir fazer utilizado a conversão do tipo inteiro para o tipo Float64
//pesopizza = pesopizza + peso
//pesopizza = float64(peso)
//pesopizza = 1 + peso