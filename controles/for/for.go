package main

import (
	"fmt"
	"time"
)

func main() {
	//exemplo com contador externo
	i := 1
	for i <= 10 { //substitui o while... já não tem while em Go
		fmt.Println(i)
		i++
	}
	fmt.Println("======================================")

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}
	fmt.Println("======================================")

	for {
		// CUIDADO isso é um loop infinito
		fmt.Println("Eternamente")
		time.Sleep(time.Second * 5)
	}
}
