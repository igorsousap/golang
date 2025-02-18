package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("olar mundo"), escrever("prgramndo go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSainda := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalDeSainda <- mensagem
			case mensagem := <-canalEntrada2:
				canalDeSainda <- mensagem
			}
		}
	}()

	return canalDeSainda
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
