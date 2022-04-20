package main

import "fmt"

// pseudoHeranca nada mas é que uma composição
// prefira composição ao invés de herança

type carro struct {
	nome            string
	velocidadeAtual uint
}

type ferrari struct {
	carro       //campo anônimo cujo tipo é um struct carro... ISSO DEFINE A COMPOSIÇÃO
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Println(f)

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
}
