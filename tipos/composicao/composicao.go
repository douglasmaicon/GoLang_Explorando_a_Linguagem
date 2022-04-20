package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	balizaAutomatica()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso

	//pode adicionar mais métodos
	autonomo()
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) balizaAutomatica() {
	fmt.Println("Baliza Automatica...")
}

func (b bmw7) autonomo() {
	fmt.Println("Dirigindo no piloto automático...")
}

func main() {
	var b1 esportivoLuxuoso = bmw7{}
	b1.ligarTurbo()
	b1.balizaAutomatica()
	b1.autonomo()
}
