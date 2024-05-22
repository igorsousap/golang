package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	go escrever("OI")
	escrever("Mundo")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
