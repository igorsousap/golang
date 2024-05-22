package main

import "fmt"

func main() {
	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	variavel5, variavel6 := "variavel5", "variavel6"

	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "Constante1"
	fmt.Println(constante1)
}
