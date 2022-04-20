package main

import "fmt"

type nota float64

func (n nota) entreConceitos(inicio, fim float64) bool {
	//return float64(n) >= inicio && float64(n) <= fim
	return n >= nota(inicio) && n <= nota(fim)
}

func notaParaConceito(n nota) string {
	if n.entreConceitos(9.0, 10) {
		return "A"
	} else if n.entreConceitos(7.0, 8.99) {
		return "B"
	} else if n.entreConceitos(5, 6.99) {
		return "C"
	} else if n.entreConceitos(3, 4.99) {
		return "D"
	} else {
		return "E"
	}

}

func main() {
	fmt.Println(notaParaConceito(9.9))
}
