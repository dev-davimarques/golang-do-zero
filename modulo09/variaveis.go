package main

import "fmt"

func main() {
	var nome string
	var idade int
	idade = 24
	nome = "davi"
	fmt.Println("Primeira atribuição: ",nome)
	nome = "marques"
	fmt.Println("Segunda atriuição: ",nome)
	fmt.Println(idade)

	// outra fomra de atribuir uma variável
	// a linguagem entende o tipo utilizado
	a := "apple"
	fmt.Println(a)

	// declarando constantes
	const sobrenome = "marques";
	nome = "davinovamente"
	fmt.Println(nome+" "+sobrenome)
}