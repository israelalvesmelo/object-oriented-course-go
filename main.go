package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	// var conta ContaCorrente = ContaCorrente{titular: "Felipe", numeroAgencia: 2343, numeroConta: 232, saldo: 233.4}
	// var conta *ContaCorrente = new(ContaCorrente)

	// conta := contas.ContaCorrente{Titular: "Felipe", NumeroAgencia: 2343, NumeroConta: 232, Saldo: 233.4}
	// conta2 := contas.ContaCorrente{Titular: "Carlos", NumeroAgencia: 2343, NumeroConta: 147, Saldo: 100.6}

	// fmt.Println(conta.MostraSaldo())

	// fmt.Println(conta.Sacar(50.), conta.Saldo)
	// fmt.Println(conta.Depositar(50.), conta.Saldo)
	// fmt.Println(conta.Transferir(150, &conta2), conta.Saldo) //Para alterar um valor de algum paramentro enviado, é necessario add o & para indicar o endereço de memoria

	// fmt.Println(conta)
	// fmt.Println(conta2)

	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "333.333.222.33",
		Profissao: "Desenvolvedor"}

	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 2343, NumeroConta: 147}
	contaDoBruno.Depositar(344.)
		PagarBoleto(&contaDoBruno, 32.)
	fmt.Println(contaDoBruno)

}

func PagarBoleto(conta verificarConta, valoDoBoleto float64){
	conta.Sacar(valoDoBoleto)
}

type verificarConta interface{
	Sacar(valor float64) string
}
