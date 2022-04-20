package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	//ao declarar uma variável do tipo interface é possivel atribuir diferentes valores nela...
	// por exemplo: coisa = 123  ou  coisa = "texto"  ou coisa = false
	// é como se fosse um tipo genérico que todos os outros herdam dele
	// funciona como um "object do java"

	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Texto qualquer"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	//inclusive é possivel usar um struct
	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)

}
