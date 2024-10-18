package main

import (
	"fmt"
)

func main() {
	// tipo booleano (true/false)
	fmt.Printf("Type %T - Valor: %v\n", true, true);
	// tipo string -> sequência de bytes
	fmt.Printf("Type %T - Valor: %v\n", "davi", "davi marques");
	// tipo inteiro
	fmt.Printf("Type %T - Valor: %v\n", 1, 1);
	// tipo inteiro ou string?
	fmt.Printf("Type %T - Valor: %v\n", "1","1");
	// tipo float(float64/float32)
	fmt.Printf("Type %T - Valor: %v\n", 1.55, 1.55);
}