package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func calcula(t Triangulo) float64 {
	area := (t.base * t.altura) / 2
	return area
}

func main() {
	r := Triangulo{base: 5, altura: 12}
	s := calcula(r)
	fmt.Println(s)
}
