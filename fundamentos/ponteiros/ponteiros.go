package main

import "fmt"

func main() {
	i := 1 //ocupa 8 bytes na memória

	// Go não tem aritmética de ponteiros mas é possível compartilhar esse ponteiro com outras referências
	var p *int = nil //ponteiro alocado mas não referencia nenhum endereço de memória
	p = &i           // & + var_name = pegando o endereço da variável
	//p++ // isso não funciona pois p não possui um valor inteiro mas na verdade o endereço de uma variável
	*p++ // isso funciona pois está incrementado o que está no endereço armazenado em p
	i++  // tbm incrementa o mesmo endereço armazenado em p

	fmt.Println(&i) //exibirá o endereço de memória da variável i
	fmt.Println(p)  //exibirá o endereço de memória de p
	fmt.Println(*p) //exibirá o conteúdo do que está sendo referenciado pelo enderço de p
	fmt.Println(i)  //exibirá o mesmo conteúdo da linha acima

}
