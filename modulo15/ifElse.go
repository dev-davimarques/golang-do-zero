package main

import(
	"fmt"
)

func main(){
	var n1 float64 = 1
	var n2 float64 = 9
	var media float64 = (n1+n2)/2
	if media >= 6 {
		fmt.Printf("%g >= 6 \n",media)
	} else{
		fmt.Printf("%g <= 6 \n",media)
	}

	// numero par ou impar
	numero := 7
	if numero%2==0 {
		fmt.Printf("%d é par \n",numero)
	} else {
		fmt.Printf("%d é ímpar \n",numero)
	}
}