package main

import "fmt"

func f1() {
	fmt.Println("Função 1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("Função 2: %s %s\n", p1, p2)
}

func f3() string {
	return "Função 3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("Função 4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno A", "Retorno B"
}

func main() {
	f1()
	f2("Douglas", "Maicon")
	r3, r4 := f3(), f4("Maria", "Luiza")
	fmt.Println(r3)
	fmt.Println(r4)
	r5a, r5b := f5()
	fmt.Println("Função 5: ", r5a, r5b)
}
