package main

import(
	"fmt"
)

func main(){
	fmt.Println(somaNumeros(3))
}

func somaNumeros(numero int)int{
	// soma := 0
	// for i := 1; i <= numero; i++ {
	// 	soma += i
	// }
	// return soma
	return numero * (numero + 1) /2
}