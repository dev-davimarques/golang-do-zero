package main

import(
	"fmt"
)

func main(){
	// Maps: Estrutras Heterogêneos

	// Criar vazio e atribuir os valores
	idade := map[string]int{}
	idade["davi"] = 24
	idade["diego"] = 30
	fmt.Println(idade)
	fmt.Println(idade["davi"])

	// Criar com os dados já preenchidos
	anoNasc := map[string]int{
		"davi": 2000,
		"diego": 1992,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["diego"])

	anoNasc["golang do zero"] = 2024
	fmt.Println(anoNasc)
}