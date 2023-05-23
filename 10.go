package main

import "fmt"

type Filme struct {
	título    string
	diretor   string
	ano       int
	avaliação []int
}

func add(f Filme, nova int) Filme {
	f.avaliação = append(f.avaliação, nova)
	return f
}

func medi(f Filme) int {
	var soma int
	for i := 0; i < len(f.avaliação); i++ {
		soma += f.avaliação[i]
	}
	med := soma / int(len(f.avaliação))
	return med
}

func main() {
	var nova int
	fmt.Print("Insira uma nova avaliação: ")
	fmt.Scanln(&nova)
	f := Filme{título: "Vingadores", diretor: "Irmãos Russos", ano: 2019, avaliação: []int{4, 5, 3}}
	r := add(f, nova)
	fmt.Println(r)
	r2 := medi(f)
	fmt.Println("A média das notas do filme é: ", r2)
}
