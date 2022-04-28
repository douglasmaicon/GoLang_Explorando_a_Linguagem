package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ { //for sem condição até achar um break
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) //fecha o canal para que o laço for não espere nenhum dado... sem isso vai gerar um deadLock
	//é uma boa prática fechar o canal para se evitar o deadlock
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c) //cap(c) representa o tamanho definido para o buffer do canal (30)

	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Terminou")
}
