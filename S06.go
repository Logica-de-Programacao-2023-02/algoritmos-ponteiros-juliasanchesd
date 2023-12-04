package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func AlterarTitulo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro1",
		Autor:  "Anônimo",
	}

	fmt.Println("Livro antes:", livro)

	AlterarTitulo(&livro)

	fmt.Println("Livro depois:", livro)
}
