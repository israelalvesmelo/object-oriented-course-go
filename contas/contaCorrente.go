package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com  sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito <= 0 {
		return "Valor para deposito invalido"
	}

	c.saldo += valorDoDeposito
	return "Valor depositado com sucesso"
}

func (c *ContaCorrente) MostraSaldo() string {
	return "Valor do saldo: " + fmt.Sprintf("%v", c.saldo)
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia > c.saldo {
		return "Valor insuficiente na conta"

	}

	if valorDaTransferencia < 0 {
		return "Valor invalido para transferencia"

	}

	c.Sacar(valorDaTransferencia)
	contaDestino.Depositar(valorDaTransferencia)
	return "Valor transferido com sucesso"

}
