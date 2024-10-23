package main

import(
	"fmt"
)

func main (){
	fmt.Println(retornaNegativo(-2))
}

func retornaNegativo(numero int)int{
	// if numero < 0{
	// 	return numero
	// } else {
	// 	return numero*-1
	// }
	if numero >= 0{
		return -numero
	}
	return numero
}