package main

import(
	"fmt"
)

func main(){
	// make recebe o tipo []string e a capacidade inicial do slice
	// cria um slice com tamanho 3
	animes := make([]string, 3)
	animes[0] = "Gantz"
	animes[1] = "Berserk"
	animes[2] = "Attack on titan"
	fmt.Println(animes)
	fmt.Println(len(animes),cap(animes))
	
	// a função append redimensiona a capacidade do slice
	animes = append(animes, "Mob Pyscho")
	fmt.Println(animes)
	fmt.Println(len(animes), cap(animes))
	for i := 0; i < len(animes); i++ {
		fmt.Println(i," - ",animes[i])
	}
	// array[inicio:fim]
	fmt.Println(animes[1:3])
}