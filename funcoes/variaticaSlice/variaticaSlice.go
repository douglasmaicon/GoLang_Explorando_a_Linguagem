package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")

	for index, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", index+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Douglas"}
	imprimirAprovados(aprovados...)
}
