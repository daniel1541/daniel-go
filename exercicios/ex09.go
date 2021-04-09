// Você possui uma turma com 4 alunos cujos nomes são Pedro, Ricardo, Paula e Marcia.
// Cada um fez 3 exames e tiveram notas diversas, porém a média final de cada deverá ser igual ou maior a 7.0 para que sejam aprovados.
// As notas de cada um deles são as seguintes:
// - Pedro: 4.5, 7.9 e 8.1
// - Ricardo: 5.7, 6.0 e 9.1
// - Paula: 6.8, 7.6 e 8.7
// - Marcia: 7.5, 5.9 e 7.6

// a. Crie uma variável que receba o nome de cada aluno. Estas variáveis devem possuir escopo geral.

// b. Para cada nota de cada aluno, crie uma variável. Estas variáveis de nota precisam ser com escopo de função.

// c. Em sua função main imprima:
// c.1- "Os alunos da turma são:"
// c.2- Na linha seguinte imprima o valor das variáveis com os nomes dos 4 alunos.
// c.3- Chame uma nova função chamada media

// d. a função media deverá imprimir:
// d.1- A média final de cada aluno indicando o nome do mesmo;
// d.2- Um resultado booleano que represente "true" para cada aluno que foi aprovado (cuja média seja igual ou maior a 7.0)

// e. Crie uma nova função chamada diplomas
// e.1- A função main deverá chamar também a função diplomas.
// e.2- A função diplomas irá imprimir a seguinte linha para cada alunoq ue foi aprovado:
// "Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana <NOME>."

// f. Crie uma função chamada reprovados.
// f.1- A função main também deverá chamar a função reprovados
// f.2- A função reprovados deverá imprimir o seguinte para cada aluno que foi reprovado:
// "Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima <ALUNO>."

package main

import (
	"fmt"
)

var nome_1 string = "Pedro"
var nome_2 string = "Ricardo"
var nome_3 string = "Paula"
var nome_4 string = "Marcia"

// - Pedro: 4.5, 7.9 e 8.1
var nota_1_nome_1 float64 = 4.5
var nota_2_nome_1 float64 = 7.9
var nota_3_nome_1 float64 = 8.1

// - Ricardo: 5.7, 6.0 e 9.1
var nota_1_nome_2 float64 = 5.7
var nota_2_nome_2 float64 = 6.0
var nota_3_nome_2 float64 = 9.1

// - Paula: 6.8, 7.6 e 8.7
var nota_1_nome_3 float64 = 6.8
var nota_2_nome_3 float64 = 7.6
var nota_3_nome_3 float64 = 8.7

// - Marcia: 7.5, 5.9 e 7.6
var nota_1_nome_4 float64 = 7.5
var nota_2_nome_4 float64 = 5.9
var nota_3_nome_4 float64 = 7.6

func main() {

	fmt.Println("Os alunos da turma são: ")
	fmt.Println(nome_1, ",", nome_2, ",", nome_3, "e", nome_4)

	media()
	diplomas()
	reporvados()
}

func media() {

	var media_nome_1 float64 = (nota_1_nome_1 + nota_2_nome_1 + nota_3_nome_1) / 3
	var media_nome_2 float64 = (nota_1_nome_2 + nota_2_nome_2 + nota_3_nome_2) / 3
	var media_nome_3 float64 = (nota_1_nome_3 + nota_2_nome_3 + nota_3_nome_3) / 3
	var media_nome_4 float64 = (nota_1_nome_4 + nota_2_nome_4 + nota_3_nome_4) / 3

	var resultado_final_1 bool = (media_nome_1 >= 7)
	var resultado_final_2 bool = (media_nome_2 >= 7)
	var resultado_final_3 bool = (media_nome_3 >= 7)
	var resultado_final_4 bool = (media_nome_4 >= 7)

	fmt.Println("Aluno(a):", nome_1, "-> Resultado se aprovado: ", resultado_final_1)
	fmt.Println("Aluno(a):", nome_2, "-> Resultado se aprovado: ", resultado_final_2)
	fmt.Println("Aluno(a):", nome_3, "-> Resultado se aprovado: ", resultado_final_3)
	fmt.Println("Aluno(a):", nome_4, "-> Resultado se aprovado: ", resultado_final_4)
}

func diplomas() {

	fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana:", nome_3)
	fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana:", nome_4)
}

func reporvados() {

	fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima aluno:", nome_1)
	fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima aluno:", nome_2)
}
