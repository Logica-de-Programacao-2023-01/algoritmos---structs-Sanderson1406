package main

import "fmt"

type Filme struct {
	título    string
	diretor   string
	ano       int
	avaliação []int
}

func remover(f Filme, nova int) Filme {
	f.avaliação = append(f.avaliação, nova)
	return f
}
func main() {
	var nova int
	fmt.Print("Insira uma nova avaliação: ")
	fmt.Scanln(&nova)

}
