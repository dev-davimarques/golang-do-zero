package main

import(
	"fmt"
)

func main(){
	// O array vai te possibilitar ter vários valores de um mesmo tipo
	var animes [3]string
	animes[0] = "Gantz"
	animes[1] = "Berserk"
	animes[2] = "Attack on titan"
	for i := 0; i < len(animes); i++{
		fmt.Println(i," - ",animes[i])
	}
	fmt.Println("-------------")

	// também pode utilizar o operador ':=' e passar os valores entre as chaves
	personagens := [3]string{"Kurono", "Gutz", "Eren"}
	for i := 0; i < len(personagens); i++{
		fmt.Println(i," - ",personagens[i])
	}
	fmt.Println("-------------")

	// ou substituindo o tamanho por '...'
	personagens2 := [...]string{"Kei Kishimoto -","Casca -","Mikasa"}
	fmt.Println(personagens2)
}