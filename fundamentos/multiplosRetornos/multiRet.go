package main

import "fmt"

func multReturnos(a, b int) (int, string, bool) {
	retInt := a + b
	retString := "cara que doidera"
	retBool := a > b
	return retInt, retString, retBool
}

func main() {
	retornoInt, retornoString, retornoBool := multReturnos(17, 14)
	fmt.Printf("Retorno Int: %v \nRetorno String: %s \nRetorno Bool: %t", retornoInt, retornoString, retornoBool)
}
