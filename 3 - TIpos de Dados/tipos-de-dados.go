package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int8 = 100
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	var numero3 int64 = -10000000
	fmt.Println(numero3)

	//int32 = RUNE
	var numero4 rune = 123456
	fmt.Println(numero4)

	//int8 = byte
	var numero5 byte = 123
	fmt.Println(numero5)

	var numero6 float32 = 123.45
	fmt.Println(numero6)

	var numero7 float64 = 12315646.45
	fmt.Println(numero7)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	var int int64
	fmt.Println(int)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var erro error = errors.New("error_interno")
	fmt.Println(erro)
}
