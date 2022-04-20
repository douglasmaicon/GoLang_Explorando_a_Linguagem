package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.141516
	var raio = 3.2

	//forma reduziada
	area := PI * math.Pow(raio, 2)

	fmt.Println("Área:", area)

	const (
		A = 1
		B = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(A, B, c, d)
}
