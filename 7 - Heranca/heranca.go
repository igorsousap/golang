package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("heranca")

	p1 := pessoa{"joao", "pedro", 20, 178}
	fmt.Println(p1)

	el := estudante{p1, "Engnharia", "USP"}
	fmt.Println(el.altura)
}
