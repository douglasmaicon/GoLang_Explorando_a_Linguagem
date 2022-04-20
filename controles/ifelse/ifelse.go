package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota: ", nota)
	} else {
		fmt.Println("Reprovado com nota: ", nota)
	}
}

func notaParaConceito(n float64) string {
	if n >= 9 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func switchNotaConceito(n float64) string {
	i := int(n)
	switch {
	case i >= 9:
		return "A"
	case i >= 8 && i < 9:
		return "B"
	case i >= 5 && i < 8:
		return "C"
	case i >= 3 && i < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)

	fmt.Println("Conceito: ", notaParaConceito(100))
	fmt.Println("Conceito: ", switchNotaConceito(100))
}
