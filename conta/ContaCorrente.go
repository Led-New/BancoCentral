package contas

import "./clientes"

type ContaCorrente struct {
    Titular       clientes.Titular
    NumeroAgencia  int
    NumeroConta    int
    Saldo          float64
}