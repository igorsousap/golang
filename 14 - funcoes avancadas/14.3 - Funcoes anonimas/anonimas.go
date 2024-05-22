package main

import "fmt"

func main() {

	func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Ola Mundo")
}
