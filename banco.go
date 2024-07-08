package main

import (
	"fmt"

	"BancoCentral/clientes"
	"BancoCentral/conta"
)

func main() {
    clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "123.123.123.12", Profissao: "Desenvolvedor"}
    contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}
    fmt.Println(contaDoBruno)
}