package main

import "fmt"

//    ...TIPO do parâmetro faz com que a função receba uma quanttidade variática no parametro
func media(numeros ...float64) float64 {
	total := 0.0

	// fazendo um range para realizar um destruct em números mas ignorando o INDEX de cada elemento
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Media: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
