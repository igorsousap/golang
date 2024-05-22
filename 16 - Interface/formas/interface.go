package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func calculaArea(f forma) {
	fmt.Printf("A area da forma e %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	calculaArea(r)

	c := circulo{10}
	calculaArea(c)
}
