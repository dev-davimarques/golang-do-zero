package main

import (
	// fmt = format
	"fmt"
)

func main(){
	var (
		nome = "Davi Bezerra Marques"
		altura = 1.85
		idade = 24
	)

	var (
		nota1 = 7.5
		nota2 = 5
	)

	var minhaMedia = imprimirMedia(float64(nota1), float64(nota2))

	imprimirCaracteristicas(nome, altura, idade)
	fmt.Println(minhaMedia)
}

func imprimirCaracteristicas(nome string, altura float64, idade int){
	fmt.Println("Seu nome é ",nome)
	fmt.Println("Sua altura é: ",altura)
	fmt.Println("Sua idade é: ",idade)
}

func imprimirMedia(nota1 float64, nota2 float64) float64{
	// var mediaAluno float64 = (nota1 + nota2) / 2
	return (nota1 + nota2) / 2
}
