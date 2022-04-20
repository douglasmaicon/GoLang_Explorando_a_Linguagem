package main

import "fmt"

type itemStruct struct {
	produtoID uint
	qtde      uint
	preco     float64
}

type pedidoStruct struct {
	userID uint
	itens  []itemStruct //slice de itens
}

func (p pedidoStruct) valorTotal() float64 {
	total := 0.0

	//percorrer a lista de itens do pedido
	fmt.Println("============  ITENS DO PEDIDO  ============")
	fmt.Println("Item| Código| Qtde| Preço")
	for idx, item := range p.itens {
		total += item.preco * float64(item.qtde)
		fmt.Printf("%d      %d      %d     %.2f\n", idx, item.produtoID, item.qtde, item.preco)
	}
	return total
}

func main() {
	pedido := pedidoStruct{
		userID: 1,
		//itens é um slice de itemStruct{}
		itens: []itemStruct{
			itemStruct{1, 2, 12.10}, //o compilador alerta sobre uma redundancia... não precisa repetir itemStruct{}
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}
	fmt.Printf("O valor total do pedido é %.2f", pedido.valorTotal())
}
