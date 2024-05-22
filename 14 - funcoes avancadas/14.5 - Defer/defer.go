package main

import "fmt"

func func1() {
	fmt.Println("Executa 1")
}

func func2() {
	fmt.Println("Executa 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Medica Calculada")
	fmt.Println("Entrando na funcao para verificar aluno aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	}

	return false
}

func main() {
	defer func1()
	func2()

	fmt.Println(alunoAprovado(7, 8))
}
