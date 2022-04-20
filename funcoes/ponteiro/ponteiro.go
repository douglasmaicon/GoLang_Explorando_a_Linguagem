package main

import "fmt"

func inc1(n int) {
	n++
	fmt.Println("O valor de n dentro da função é incrementado: ", n)
	fmt.Println("Mas fora da função não permanece o mesmo.")
}

//um ponteiro é representado por um *
func inc2(n *int) {
	//* tbm é usado para obter o valor ao qual o ponteiro está referenciado
	*n++
	fmt.Println("Como foi passado um ponteiro *n, o valor é alterado fora da função")
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n) //

	//& é usado para obter o endereço em memória de uma variável (para o ponteiro)
	inc2(&n)
	fmt.Println(n)
}
