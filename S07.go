package main

type Conta struct {
	Saldo   float64
	Titular string
}

func AdicionarSaldo(conta *Conta, valor float64) {
	conta.Saldo += valor
}
