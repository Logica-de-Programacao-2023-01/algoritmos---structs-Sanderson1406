package main

import (
	"fmt"
	"time"
)

type Playlist struct {
	nome   string
	musica []Musica
}

type Musica struct {
	titulo  string
	artista string
	duracao time.Duration
}

func nome(p Playlist) string {
	return p.nome
}

func infotempo(p Playlist) {
	var tempo time.Duration
	for _, ms := range p.musica {
		tempo += ms.duracao
		fmt.Println("TITULO", ms.titulo)
		fmt.Println("ARTISTA", ms.artista)
		fmt.Println("DURAÇÃO", ms.duracao)
	}
	fmt.Println("Tempo Total: ", tempo)
}

func main() {
	p := Playlist{
		nome: "Top 10",
		musica: []Musica{
			{
				titulo:  "Vamos la",
				artista: "Marilia",
				duracao: 3*time.Minute + 40*time.Second,
			},
			{
				titulo:  "Oi, tvhau",
				artista: "Sand",
				duracao: 2*time.Minute + 33*time.Second,
			},
			{
				titulo:  "Vomo lindos",
				artista: "Pedro",
				duracao: 4*time.Minute + 22*time.Second,
			},
			{
				titulo:  "Pq assimw",
				artista: "Juliana",
				duracao: 3*time.Minute + 12*time.Second,
			},
		},
	}
	r1 := nome(p)
	fmt.Println(r1)
	infotempo(p)
}
