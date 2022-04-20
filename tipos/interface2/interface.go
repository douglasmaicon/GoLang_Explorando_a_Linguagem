package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// lembrando que é preciso trabalhar com ponteiro para garantir que todos os objetos do tipo ferrari struct
// tenham acesso ao método ligarTurbo
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	ferr := ferrari{"F40", false, 0} //trabalhando diretamente com o tipo ferrari{}
	ferr.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0} // trabalhando com a interface é preciso usar ponteiro
	carro2.ligarTurbo()

	fmt.Println(ferr, carro2)
}
