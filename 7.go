package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func mudarson(a Animal, novoSom string) Animal {
	a.som = novoSom
	return a
}

func main() {
	novoSom := "MUUU"
	a := Animal{
		nome:    "Lindi",
		especie: "Gato",
		idade:   8,
		som:     "Miau",
	}
	r := mudarson(a, novoSom)
	fmt.Println(r)
}
