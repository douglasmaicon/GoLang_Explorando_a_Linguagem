package main

import "fmt"

func trocar(p1, p2 int) (primeiro, segundo int) {
	segundo = p1
	primeiro = p2
	return //retorno limpo, já está atribuído ao retorno nomeado
}

func main() {
	r1, r2 := trocar(10, 20)
	fmt.Println(r1, r2)
}
