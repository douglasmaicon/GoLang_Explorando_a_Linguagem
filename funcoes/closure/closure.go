package main

import "fmt"

//declaração de uma função closure quem retorna uma outra função
func closure() func() {
	x := 10

	// atribuindo uma função anônima à variavel funcao
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()
}
