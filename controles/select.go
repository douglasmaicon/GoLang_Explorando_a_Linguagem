package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	c1 := make(chan string) //cria um canal string
	c2 := make(chan string) //cria um canal string

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one" //envia "one" para o canal c1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two" // envia "two" para o canal c2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: /// duvida: msg1 é um canal? (não foi usado make)
			fmt.Println(reflect.TypeOf(msg1)) // o tipo retornado é string
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
