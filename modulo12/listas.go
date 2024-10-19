package main

import (
	"fmt"
)

func main(){
	// Array de tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2,3,5,7,11,13}
	fmt.Println(numPrimos)
	// printar posicoes especificas
	fmt.Println(numPrimos[0:3])
	// printar somente um posicao
	fmt.Println(numPrimos[5])
	// printar algo que está antes de uma posição, ex: tudo que está antes da posicao 3
	fmt.Println(numPrimos[:3])
	// printar algo que está depois de uma posição, ex: tudo que está depois da posicao 3
	fmt.Println(numPrimos[3:])


	// Slice não possui tamanho definido
	// a funcao make serve para criar o slice zerado
	slice := make([]string, 5)
	slice[0] = "Davi"
	slice[1] = "Marques"
	fmt.Println(slice[0],slice[1])
	fmt.Println(slice)
	slice[2] = "Bezerra"
	fmt.Println(slice[0],slice[2],slice[1])
	slice[0] = "Davizera"
	fmt.Println(slice[0],slice[2],slice[1])
	// função para retornar tamanho do slice
	fmt.Println(len(slice))

	// caso você não queira criar o slice zerado
	numPares := []int{2,4,6,8,10}
	fmt.Println(numPares)
	fmt.Println(len(numPares))
	// funcao append adiciona valores no slice
	numPares = append(numPares, 12)
	fmt.Println(numPares)
	fmt.Println(len(numPares))
	// podemos adicionar vários itens de uma só vez
	numPares = append(numPares, 14,16,18,20)
	fmt.Println(numPares)
	fmt.Println(len(numPares))
}