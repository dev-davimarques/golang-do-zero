package main

import(
	"fmt"
)

func main(){
	var soma int;
	for i := 0; i < 10; i++ {
		soma += i
		fmt.Printf("[%d] - %d \n",i,soma)
	}

	// for range
	frutas := []string{"laranja", "kiwi", "banana", "uva", "maçã"}
	for i,fruta := range frutas {
		fmt.Printf("Índice [%d] - %s \n",i,fruta)
	}
}