package main

import (
	"Banco/conta"
	"fmt"
)
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
    conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
    Sacar(valor float64) string
}

func main() {
    contaDoDenis := contas.ContaPoupança{}
    contaDoDenis.Depositar(100)
    PagarBoleto(&contaDoDenis, 19)

    fmt.Println(contaDoDenis.SaldoConta())
}
