package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO REFERENCIA DE MEMORIA

	var variavel3 int = 100
	var ponteiro *int

	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)

}
