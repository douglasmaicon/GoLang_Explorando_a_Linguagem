package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(array1,4,5,6) o append exige que o primeiro argumento seja um slice1

	slice1 = append(slice1, 4, 5, 6) //se o slice já estiver com a capacidade máxima irá duplicar o array interno de referência
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	//copy não expande o slice como ocorre com o append
	//copy exige que o destino e a origem sejam slices
	copy(slice2, array1[:]) //array1[:] faz um slice completo do array1
	fmt.Println(slice2)
}
