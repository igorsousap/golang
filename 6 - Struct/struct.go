package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo struct")

	var u usuario
	u.nome = "Igor"
	u.idade = 28
	fmt.Println(u)

	endereco_ex := endereco{"Rua teste", 0}

	usuario2 := usuario{"igor", 28, endereco_ex}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "igor"}
	fmt.Println(usuario3)
}
