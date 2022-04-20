package main

import "fmt"

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("As funções init() são executadas antes da função main()")
}
