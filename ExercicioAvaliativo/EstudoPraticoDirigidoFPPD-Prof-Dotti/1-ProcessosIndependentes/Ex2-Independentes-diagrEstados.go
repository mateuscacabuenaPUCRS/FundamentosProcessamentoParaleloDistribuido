// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// EXERCÍCIO:  dado o programa abaixo
//     considerando-se como estados os valores da tripla x,y,z
//     qual o diagrama de estados e transicoes que representa
//     1) a questaoStSp2()  z permanece 0, x é 1 enquanto y é 0, se tornar 2 e dps o y vira 2.
//     2) a questaoStSp3()  x e y podem se tornar 1 ou 2 antes dos mesmos alterarem pq é uma concorrência e, por ultimo, z vira 2.
// OBS: a execucao do programa abaixo nao mostra nada.   este serve como especificacao do problema.
//      note que como não há sincronizacao, todas combinacoes possiveis de estados acontecerao.

package main

import "fmt"

//---------------------------

var x, y, z int = 0, 0, 0

func px() {
	x = 1
	x = 2
}

func py() {
	y = 1
	y = 2
}

func pz() {
	z = 1
	z = 2
}

func questaoStSp2() {
	go px()
	py()
	for {
	}
}

func questaoStSp3() {
	go px()
	go py()
	pz()
	for {
		fmt.Println(x, y, z)
	}
}

func main() {
	questaoStSp2()
	questaoStSp3()
}
