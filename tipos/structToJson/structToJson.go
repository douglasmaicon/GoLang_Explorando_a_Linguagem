package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    uint     `json:"id"` //primeira letra maiuscula indica que será um atributo público
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.99, []string{"Promoção", "Eletrônicos"}}
	//desestruturação de json.Marshal ignorando o error (2ª retorno)
	p1Json, _ := json.Marshal(p1)
	fmt.Println((string(p1Json)))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Mouse","preco":45.99,"tags":["Eletrônicos","Periféricos"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
