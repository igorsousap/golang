package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal2 <- "canal2"
		}
	}()

	for {
		select {
		case mensagemCanil1 := <-canal1:
			fmt.Println(mensagemCanil1)
		case mensagemCanil2 := <-canal2:
			fmt.Println(mensagemCanil2)
		}
	}

}
