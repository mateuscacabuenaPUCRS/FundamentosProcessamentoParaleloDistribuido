// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// PROBLEMA:
//   o dorminhoco especificado no arquivo Ex1-ExplanacaoDoDorminhoco.pdf nesta pasta
// ESTE ARQUIVO
//   Um template para criar um anel generico.
//   Adapte para o problema do dorminhoco.
//   Nada está dito sobre como funciona a ordem de processos que batem.
//   O ultimo leva a rolhada ...
//   ESTE  PROGRAMA NAO FUNCIONA.    É UM RASCUNHO COM DICAS.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NJ = 5 // numero de jogadores
const M = 4  // numero de cartas

type carta string // carta é um strirng

var ch [NJ]chan carta // NJ canais de itens tipo carta

var bateu [NJ]chan int

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta) {
	mao := cartasIniciais // estado local - as cartas na mao do jogador
	nroDeCartas := M      // quantas cartas ele tem
	cartaRecebida := carta("")
	fmt.Printf("Jogador %d cartas na mao: %v\n", id, mao)

	for {
		if nroDeCartas == 5 {
			fmt.Println(id, " joga") // escreve seu identificador
			// escolhe alguma carta para passar adiante ...
			indiceAleatorio := rand.Intn(nroDeCartas)
			cartaParaSair := mao[indiceAleatorio]

			// manda carta escolhida o proximo
			out <- cartaParaSair
			mao = append(mao[:indiceAleatorio], mao[indiceAleatorio+1:]...)
			nroDeCartas--
		} else {
			cartaRecebida = <-in
			fmt.Printf("Jogador %d recebeu a carta %s\n", id, cartaRecebida)
			mao = append(mao, cartaRecebida)
			nroDeCartas++

			// Verifica se um jogador possui 4 cartas da mesma letra e bate
			if possuiQuatroLetrasIguais(mao) {
				fmt.Printf("Jogador %d bate!\n", id)
				//bateu[id] <- 1
				return;
			}
		}
	}
}

func possuiQuatroLetrasIguais(mao []carta) bool {
	contagem := make(map[carta]int)
	for _, c := range mao {
		contagem[c]++
		if contagem[c] >= 4 {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < NJ; i++ {
		ch[i] = make(chan carta)
	}

	for i := 0; i < NJ; i++ {
		cartasEscolhidas := make([]carta, 4)
		for j := 0; j < 4; j++ { // escolhe aleatoriamente (tira) cartas do baralho, passa cartas para jogador
			indiceAleatorio := rand.Intn(4) // Gera aleatoriamente um índice entre 0 e 3
			cartasEscolhidas[j] = carta('A' + indiceAleatorio)
		}

		go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas)
	}

	// Comece o jogo passando uma carta inicial
	cartaInicial := carta("@")
	ch[0] <- cartaInicial

	// Aguarde todos os jogadores terminarem
	for i := 0; i < NJ; i++ {
		<-ch[i]
	}

	fmt.Println("O jogo terminou!")
	<-make(chan carta) // bloqueia
}
