package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// INTERFACES SÃO IMPLEMENTADAS IMPLICITAMENTE
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Douglas", "Nascimento"}
	fmt.Println(coisa.toString()) // vai executar a func toString que tem pessoa como receiver
	imprimir(coisa)               //vai executar o toString implicitamente

	coisa = produto{"Monitor", 999.99}
	fmt.Println(coisa.toString()) // vai executar a func toString que tem produto como receiver
	imprimir(coisa)               //vai executar o toString implicitamente

	coisa2 := produto{"PlayStation 5", 4999.99} // não precisa ser do "tipo" imprimivel para funcionar, basta respeitar a interface
	fmt.Println(coisa2.toString())
}
