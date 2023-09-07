// Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// Exemplo da Internet
// EXERCICIOS:
//   1) rode o programa abaixo e interprete.
//      todos os valores escritos no canal são lidos? Não, apenas 5 dos 10.
//   2) como isto poderia ser resolvido ? Aumentando o canal para 10 inteiros invés de apenas 5.

package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	go shower(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func shower(c chan int) {
	for {
		j := <-c
		fmt.Printf("%d\n", j)
	}
}
