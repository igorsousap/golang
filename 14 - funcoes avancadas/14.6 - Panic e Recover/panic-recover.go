package main

import "fmt"

func recuperarExecute() {
	if r := recover(); r != nil {
		fmt.Println("Recupera rexecusao")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecute()
	fmt.Println("Entrando na funcao para verificar aluno aprovado")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media igual a 6")
	//kill execution
}

func main() {

	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pos execusao")
}
