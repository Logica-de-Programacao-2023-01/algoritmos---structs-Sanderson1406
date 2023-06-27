package main

import "fmt"

type Musica struct {
	Titulo  string
	Artista string
	Duracao int
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func buscarPlaylistsPorTitulo(playlists []Playlist, titulo string) []Playlist {
	var resultado []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == titulo {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	playlist1 := Playlist{
		Nome: "Playlist 1",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4},
		},
	}

	playlist2 := Playlist{
		Nome: "Playlist 2",
		Musicas: []Musica{
			{Titulo: "Música 2", Artista: "Artista 3", Duracao: 5},
			{Titulo: "Música 3", Artista: "Artista 4", Duracao: 6},
		},
	}

	playlists := []Playlist{playlist1, playlist2}

	titulo := "Música 2"
	resultado := buscarPlaylistsPorTitulo(playlists, titulo)

	fmt.Printf("Playlists com o título \"%s\":\n", titulo)
	for _, playlist := range resultado {
		fmt.Println(playlist.Nome)
	}
}
