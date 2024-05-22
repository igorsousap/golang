package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus1",
		},
	}

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome_signo": "Gemeos",
	}

	fmt.Println(usuario2)
}
