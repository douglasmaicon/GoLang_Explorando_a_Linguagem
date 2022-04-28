package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou") //aqui pode ser que ainda execute pois já pode ter lido o buffer[3] e ja tenha liberado espaço no buffer
	ch <- 5
	//fmt.Println("Executou") //aqui não vai mais executar DE FORMA NENHUMA pois não tem mais estado no buffer do chanel
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
