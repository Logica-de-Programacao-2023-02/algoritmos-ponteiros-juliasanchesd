package main

import "fmt"

func ModificarValor(ptr *int) {
	*ptr = 42
}

func main() {
	var valor int
	ponteiro := &valor

	fmt.Println("Valor antes:", valor)

	ModificarValor(ponteiro)

	fmt.Println("Valor depois:", valor)
}
