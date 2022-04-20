package main

import "fmt"

//função anônima
// não usou := pois isso só é possível dentro de uma func
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	// := para fazer atribuição por inferência
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 3))
}
