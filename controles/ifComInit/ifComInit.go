package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //esse padrão tbm é suportado pelo switch
		//a variável i somente é acessível no if ou no seu else

		fmt.Println("Ganhou: ", i)
	} else {
		fmt.Println("Perdeu: ", i)
	}
}
