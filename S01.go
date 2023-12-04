package main

import "fmt"

package Ponteiros

import "fmt"

func SomaDosNaturais(ptr *int, n int) {
	if n < 0 {
		fmt.Print("O valor de n deve ser um número não negativo.")
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	*ptr = sum
}