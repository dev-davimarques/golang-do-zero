package main

import (
	"fmt"
)

func main(){
	// Structs
	// Forma de criar sua própria estrutura de dados
	// Personalizar de acordo com a sua capacidade
	// Podemos usar vários tipos diferentes

	// type <NOME_DA_ESTRUTURA> struct { <campos da estrutura> }
	type Pessoa struct{
		Nome string
		Idade int
		// Sobrenome string
		// AnoNasc int
		// Ativo bool
	}
	type Profissao struct{
		Pessoa
		Tipo string
	}
	// o go aceita dessa forma
	fmt.Println(Pessoa{"Davi",24})
	// mas essa é a boa prática
	fmt.Println(Pessoa{Nome: "Davi", Idade: 24})
	// também é possível mostrar somente uma parte da estrutura
	fmt.Println(Pessoa{Nome: "Diego"})


	// Outra forma de usar é atribuindo esse valor a variável
	p1 := Pessoa{Nome: "Davi Marques", Idade: 24}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)
	// Posso atribuir outros valores na estrutura
	p1.Idade = 25
	fmt.Println(p1)
	p2 := Pessoa{Nome: "João", Idade: 20}

	// criando um estrutura com lista de pessoas
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// struct herdando campos de outra struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}