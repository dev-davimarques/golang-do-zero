package main

import(
	"fmt"
)

func main(){
	// são iniciados com make, é opcional informar a capacidade inicial
	// e assim é definido map[tipo da chave]tipo do valor
	dicionario := make(map[string]string)
	dicionario["hi"] = "oi"
	dicionario["bye"] = "tchau"
	dicionario["hey"] = "opa"
	fmt.Println(dicionario)

	// também podem ser iniciados com valores literais
	personagemAnime := map[string]string{
		"Eren": "Attack on Titan",
		"Guts": "Berserk",
		"Kurono": "Gantz",
	}
	fmt.Println(personagemAnime)
}