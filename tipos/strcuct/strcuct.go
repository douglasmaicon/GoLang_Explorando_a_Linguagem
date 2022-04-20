package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
	qtde     uint
}

// receiver (RECEPTOR)
//Método: função com receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
		// não é obrigatório definir todos os atributs
		//qtde:     1000,
	}
	produto2 := produto{"Monitor Dell", 799.99, 0.10, 37}

	fmt.Println("Produto1: ", produto1)
	fmt.Printf("O valor com desconto é %.2f\n", produto1.precoComDesconto())

	fmt.Println("Produto 2:", produto2)
}
