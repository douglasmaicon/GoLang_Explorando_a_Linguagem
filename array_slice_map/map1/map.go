package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem sempre inicializados

	aprovados := make(map[int]string)
	aprovados[123456] = "Maria"
	aprovados[456123] = "Douglas" //exibindo um "registro" específico
	aprovados[654321] = "Virginia"

	fmt.Println(aprovados)
	fmt.Println(aprovados[456123])

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 456123) //excluindo um "registro" específico do map
	fmt.Println(aprovados)
}
