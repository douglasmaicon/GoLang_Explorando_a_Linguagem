package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 123.32,
			"Gustavo": 555.55,
		},
		"D": {
			"Douglas": 111.11,
			"Debora":  222.22,
		},
		"V": {
			"Virginia": 777.77,
		},
	}
	delete(funcsPorLetra, "D")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println("   ", nome, salario)
		}
	}

}
