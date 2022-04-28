package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan (cria um canal somente leitura)

// um Generator é uma função que encapsula uma chamada de goRoutine

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		//cria uma função anonima
		//fmt.Println(urls)
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			//fmt.Println("situação de regex: ", r)
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.google.com.br", "https://parrotsec.org")
	t2 := titulo("https://www.cod3r.com.br", "https://youtube.com")

	//t1 := titulo("https://parrotsec.org", "https://www.google.com")
	//t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	//fmt.Println("Teste: ", <-t2, "|", <-t2)
	// o retorno da função titulo é um chanel, então t1 e t2 são canais
	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
