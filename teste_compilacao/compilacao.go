package main

import (
	"fmt"
	"time"
)

func main() {
	inicio := time.Now().UnixNano() / int64(time.Millisecond)
	a := 0
	for i := 0; i < 1000000000; i++ {
		a++
	}
	fim := time.Now().UnixNano() / int64(time.Millisecond)
	diff := fim - inicio
	fmt.Printf("Duração (ms): %d", diff)

}
