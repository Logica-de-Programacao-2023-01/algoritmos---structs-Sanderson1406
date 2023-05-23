package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

func preco(a []Viagem) Viagem {
	var max Viagem
	for _, ns := range a {
		if ns.preco > max.preco {
			max = ns
		}
	}
	return max
}

func main() {
	a := []Viagem{
		{origem: "Rio", destino: "São Paulo", data: "14/05/2022", preco: 520},
		{origem: "Maragogi", destino: "São Luis", data: "23/04/2022", preco: 600},
		{origem: "Porto Seguro", destino: "Brasília", data: "14/06/2005", preco: 650},
	}
	r := preco(a)
	fmt.Println(r)
}
