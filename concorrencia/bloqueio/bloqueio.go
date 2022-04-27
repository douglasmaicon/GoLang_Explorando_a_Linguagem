package main

import (
	"fmt"
	"time"
)

func rotina(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "Douglas" //operação bloqueante se o canal não tem um buffer
	fmt.Println("foi escrito no canal")
}

func main() {
	meuCanal := make(chan string) // canal sem buffer
	go rotina(meuCanal)
	fmt.Println((<-meuCanal)) //operação bloqueante
	fmt.Println("Foi lido do buffer")
	fmt.Println((<-meuCanal)) // aqui vai gerar o deadlock

}
