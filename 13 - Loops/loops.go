package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementamdo i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementamdo i", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Joao", "Davi", "Pedro"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Igor",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
