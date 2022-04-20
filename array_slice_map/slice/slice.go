package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //isso é um array
	s1 := []int{1, 2, 3}  // isso é um slice

	fmt.Println(a1, reflect.TypeOf(a1))
	fmt.Println(s1, reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice não é um array! mas define um pedaço de um array
	//importante saber que o slice não cria um novo array, mas
	//   tem um ponteiro para os elementos indicados na criação do slice
	s2 := a2[1:3]
	fmt.Println(a2, reflect.TypeOf(a2))
	fmt.Println(s2, reflect.TypeOf(s2))

	s3 := a2[:4]
	fmt.Println(a2, reflect.TypeOf(a2))
	fmt.Println(s3, reflect.TypeOf(s3))

	//slice de um slice
	s4 := s3[:2]
	fmt.Println(s4, reflect.TypeOf(s4))

	t := "Douglas Maicon"
	s5 := t[0:7]
	fmt.Println(s5, reflect.TypeOf(s5))
}
