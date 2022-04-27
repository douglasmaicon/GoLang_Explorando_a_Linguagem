package main

import (
	"fmt"
	"time"
)

//Chanel (canal) - é a forma de comunicação entre goRoutines
//Chanel é síncrono
//Chanel é um tipo da linguagem

func mult(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second * 3)
	c <- 4 * base
}

func main() {
	meuCanal := make(chan int)
	go mult(5, meuCanal)
	fmt.Println("Executando a goRoutine, aguardando")

	//a e b recebem os valores atribuídos ao canal pela função mult()
	a, b := <-meuCanal, <-meuCanal //só vai continuar quando tiver computado o valor de a e b (síncrono)
	fmt.Println("Agora terminou de executar parte da goRoutine para a e b")
	fmt.Println(a, b)
	fmt.Println("Aguardando executar a ultima parte da goRoutine")
	fmt.Println(<-meuCanal)
	fmt.Println("Agora terminou de executar a goRoutine")

	fmt.Println(<-meuCanal)
	// aqui vai gerar um deadlock pois não há mais nada para ser executado na goRoutine
	//com isso, não há mais nada que possa ser lido no canal
}
