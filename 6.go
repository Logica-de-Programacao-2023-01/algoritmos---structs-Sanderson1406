package main

import "fmt"

type Funcionarios struct {
	nome    string
	salario float64
	idade   int
}

func salario(f Funcionarios) float64 {
	var newsala float64 = f.salario + (f.salario * 0.9)
	return newsala
}

func tempo(f Funcionarios) int {
	var tempo int
	tempo = f.idade - 18
	return tempo
}

func main() {
	f := Funcionarios{
		nome:    "Sanderson",
		salario: 1000,
		idade:   58,
	}
	r := salario(f)
	r2 := tempo(f)
	fmt.Println("O seu novo salario é: ", r)
	fmt.Println("Seu tempo de trabalho é: ", r2)
}
