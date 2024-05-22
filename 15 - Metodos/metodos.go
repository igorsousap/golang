package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvar usaurio: %s no banco\n", u.nome)
}

func (u usuario) maiordeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"usuario 1", 20}
	fmt.Println(usuario1)

	usuario1.salvar()
	maiordeIdade := usuario1.maiordeIdade()
	fmt.Println(maiordeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
