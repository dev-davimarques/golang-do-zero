package main

import(
	"fmt"
	"strings"
)

func main(){
	fmt.Println(repeatStrings(2, "Davi"))
}

func repeatStrings (repetitions int, value string)string{
	return strings.Repeat(value, repetitions)
}