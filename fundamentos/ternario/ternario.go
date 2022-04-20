package main

import "fmt"

func obterResultado(nota float64) string {

	// se existisse o operado ternário seria algo do tipo
	// return nota >= 6 ? "Aprovado" : "Reprovado"

	if nota >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	// na verdade no go NÃO EXISTE OPERADOR TERNÁRIO
	fmt.Println(obterResultado(6.2))
}
