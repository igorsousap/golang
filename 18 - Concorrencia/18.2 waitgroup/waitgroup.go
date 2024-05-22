package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("OI")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
