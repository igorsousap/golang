package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terca-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Nao achei o final de semana"
	}

}

func diaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terca-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Nao achei o final de semana"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("switch")

	dia := diaDaSemana(3)

	fmt.Println(dia)
}
