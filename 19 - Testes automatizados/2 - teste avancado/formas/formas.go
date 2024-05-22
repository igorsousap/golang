package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func CalculaArea(f Forma) {
	fmt.Printf("A area da forma e %0.2f\n", f.area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
