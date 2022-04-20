package main

import "fmt"

func main() {
	// array são estruras homogêneas (mesmo tipo) e estático (tamanho fixo)
	var notas [3]float64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Media %.2f:\n", media)

	// for range é bem semelhante ao for each
	numeros := [...]int{1, 2, 3, 4, 5} //o compilador irá definir o tamanho do array

	// range numeros irá retornar o índice e o seu valor no array números
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	//ignorando o indice com _
	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
