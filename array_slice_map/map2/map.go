package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"douglas":  123.45,
		"maria":    456.78,
		"virginia": 789.99,
	}

	funcESalarios["uaaaaai"] = 999.99
	funcESalarios["xiqueee"] = 000.00
	delete(funcESalarios, "inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
