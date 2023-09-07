// Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// >>> Veja antes Ex4 desta série.
// EXERCICIOS:
//     1) esta é uma solução para a questão anterior ? Sim, pois acaba se cair no quit
//     2) o que garante que todos valores serão lidos antes do programa acabar ? O programa garante que todos os valores serão lidos antes de sair porque ele primeiro envia todos os valores para o canal ch e só então sinaliza a goroutine shower para parar, aguardando que ela termine de processar todos os valores antes de encerrar o programa.

package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	quit <- false // or true, does not matter
}

func shower(c chan int, quit chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}
