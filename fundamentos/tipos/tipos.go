package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)...
	//uint8  => byte (8bits)
	//uint16 => short (2 bytes)
	//uint32 => int (4 bytes)
	//uint64 => long (8 bytes	)
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i64 := math.MaxInt64
	fmt.Println("o valor máximo do int é", i64)
	fmt.Println("o tipo de i64 é", reflect.TypeOf(i64))

	//representação inteira do valor de um caracter na tabela unicode (int32)
	var iRune rune = 'a'
	fmt.Println("o valor de iRune é", iRune)
	fmt.Println("o tipo de iRune é", reflect.TypeOf(iRune))

	//numeros reais (float32 e float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literação 49.99 é", reflect.TypeOf(49.99))

	//boolena
	bot := true
	fmt.Println("o tipo de bot é", reflect.TypeOf(bot))
	fmt.Println(!bot)

	//string (somente aspas duplas ou crase)
	s1 := "olá Douglas"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho de s1 é", len(s1))

	//string com multiplas linhas (crase)
	s2 := `Olá 
	este texto
	está quebrando 
	linhas`
	fmt.Println(s2 + "!")
	fmt.Println("o tamanho de s1 é", len(s2))

	//o tipo char não existe em go
	char := "a"
	fmt.Println("o valor da variavel char é " + char)
	fmt.Println("o tipo da variavel char é", reflect.TypeOf(char))
}
