package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	nota  []float64
}

func remove(a Aluno, novaNota float64) Aluno {
	a.nota = append(a.nota, novaNota)
	return a
}

func media(a Aluno) float64 {
	var soma float64
	var resul float64
	for i := 0; i < len(a.nota); i++ {
		soma += a.nota[i]
	}
	resul = soma / float64(len(a.nota))
	return resul
}

func main() {
	var novaNota float64
	fmt.Println("Nota para add: ")
	fmt.Scanln(&novaNota)
	a := Aluno{
		nome:  "David",
		idade: 18,
		nota:  []float64{10.1, 9.8, 7.5},
	}
	r := remove(a, novaNota)
	fmt.Println(r)
	s := float64(media(a))
	fmt.Printf("A média de notas é: %2.f", s)
}
