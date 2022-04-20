package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado... strint(97) não é "97" mas sim o valor correspondente na tabela ascii "a"
	fmt.Println("teste " + string(97))

	//int para string
	fmt.Println("teste " + strconv.Itoa(97))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	fmt.Println(reflect.TypeOf(num))

	//o _ é usado nesse caso pois a função strconv.Atoi retorna 2 valores (um int e um error)
	//dessa forma o _ é usado para evitar o problema de variavel não utilizada no goLang
	// pois se fosse num, erro := strconv.Atoi("123") obrigatoriamente teria que usar a variável erro

	//detalhe: strconv tbm faz conversão de tipo boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("verdadeiro")
	} else {
		fmt.Println("falso")
	}
}
