package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Maria está falando", 3)
	//fale("João", "Só posso falar depois da Maria", 1)

	//acrescentar a palavra reservada go antes da chamada de uma função transforma essa function em uma
	//goRoutine (thread) que será executada de forma assíncrona.
	// Porém a thread principal do progrma (func main()) por terminar antes que
	//a função transformada em goRountine termine sua execução
	//go fale("Maria", "Maria está falando", 500)
	//go fale("João", "João está falando ao mesmo tempo que Maria", 500)
	//fmt.Println("func Main Terminou")

	//a mesma função pode ser executada tanto de forma assíncrona como de forma síncrona
	//nesse caso a aplicação será encerrada mesmo que a Maria não tenha falado todas as vezes
	go fale("Maria", "Maria está falando", 500)
	fale("João", "João está falando ao mesmo tempo que Maria", 5)

	//para garantir que a goRoutine seja executada por completo é necessário usar canais (chanel)
	//que será visto na proxima aula
}
