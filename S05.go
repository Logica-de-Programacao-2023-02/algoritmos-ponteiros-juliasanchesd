package main

import (
	"fmt"
	"math"
)

func AtualizarMediaComPi(ptr *float64) {
	valorAtual := *ptr
	media := (valorAtual + math.Pi) / 2
	*ptr = media
}

func main() {
	valor := 3.14
	fmt.Println("Valor antes:", valor)

	AtualizarMediaComPi(&valor)

	fmt.Println("Valor depois:", valor)
}
