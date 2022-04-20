package main

import "fmt"

func verificarIdade(idade uint) {
	defer fmt.Println("executou essa linha antes de terminar	!!!")
	defer fmt.Println("Terminou a execução dessa função!!!")

	if idade >= 18 {
		fmt.Println("Maior idade penal")
	} else {
		fmt.Println("Menor infrator, coitadinho")
	}
}

func main() {
	verificarIdade(37)
}
