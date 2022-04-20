package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

//nesse caso tem que ser por referência (ponteiro) para que o valor da variavel p1 seja alterada
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Douglas", "Maicon"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Luiza")
	fmt.Println(p1.getNomeCompleto())

}
