package main

import (
	"math"
)

type Circuito struct {
	raio float64
}

func calcular(r Circuito) float64 {
	area := r.raio * r.raio * math.Pi
	return area
}

func main() {
	r := Circuito{raio: 4}
	calcular(r)
}
