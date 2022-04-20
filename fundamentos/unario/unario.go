package main

import "fmt"

func main() {
	x := 1
	y := 2

	// go só tem a forma posfix para incremento
	x++ // x += 1  ou x = x + 1
	fmt.Println(x)

	y--
	fmt.Println(y)

	//fmt.Print(x == y--)
	// em Go não é possível esse tipo de comparação, pois tornaria o código mais complexo e de dificil leitura
	// o correto é manipular a variável antes da comparação
}
