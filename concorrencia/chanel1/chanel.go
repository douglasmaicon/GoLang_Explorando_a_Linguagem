package main

import "fmt"

func main() {
	//chanel é um tipo na linguagem go

	//sintaxe para criar um canal
	ch := make(chan int, 1)

	//enviar dados para um canal
	ch <- 1 // escrita: enviando dados para o canal
	<-ch    //leitura: recebendo dados do canal
	ch <- 2
	fmt.Println(<-ch)

	//o canal é um mecanismo síncrono, sendo assim um canal pode receber
	//uma goRoutine, que será executada por completo, mas isso pode gerar
	//uma situação de deadLock CUIDADO
}
