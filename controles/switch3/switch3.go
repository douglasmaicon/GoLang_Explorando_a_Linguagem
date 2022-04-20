package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	//o tipo de dado "interface" é um tipo genérico que permite que o método receba
	//diferentes tipos genéricos no parametro (hora é um int, hora é uma string, hora é um bool...)
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	case bool:
		return "booleano"
	default:
		return "vai saber que tipo é esse =/"
	}
}

func main() {
	fmt.Println("Tipo: ", tipo(2.5))
	fmt.Println("Tipo: ", tipo(50))
	fmt.Println("Tipo: ", tipo("Texto"))
	fmt.Println("Tipo: ", tipo(func() {}))
	fmt.Println("Tipo: ", tipo(true))
	fmt.Println("Tipo: ", tipo(time.Now()))
}
