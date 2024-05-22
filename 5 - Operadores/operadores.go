package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	sub := 1 - 2
	divisao := 10 / 4
	multip := 10 * 5
	resto_div := 10 % 2

	fmt.Println(soma, sub, divisao, multip, resto_div)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)

	fmt.Println(soma2)

	// atribuicao

	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	//operadores relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//operadores unarios
	numero11 := 10
	numero11++
	numero11 += 15
	fmt.Println(numero11)

	//operadores ternarios
	var texto string
	if numero11 > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}

	fmt.Println(texto)
}
