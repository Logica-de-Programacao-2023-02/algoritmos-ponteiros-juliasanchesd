package main

func SomaUltimosDigitos(ptr *int) {
	numero := *ptr

	digito1 := numero % 10
	numero /= 10
	digito2 := numero % 10

	*ptr = digito1 + digito2
}
