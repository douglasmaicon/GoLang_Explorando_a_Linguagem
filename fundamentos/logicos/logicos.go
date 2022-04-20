package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvet := trab1 || trab2
	return comprarTv50, comprarTv32, comprarSorvet
}

func multReturnos(a, b int) (int, string, bool) {
	retInt := a + b
	retString := "cara que doidera"
	retBool := a > b
	return retInt, retString, retBool
}

func main() {
	//tv50, tv32, sorvete := compras(true, true)
	//fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
	retornoInt, retornoString, retornoBool := multReturnos(17, 32)
	fmt.Printf("Retorno Int: %v \nRetorno String: %s \nRetorno Bool: %t", retornoInt, retornoString, retornoBool)
}
