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
)

const NJ = 5           // numero de jogadores
const M = 4            // numero de cartas

type carta string      // carta é um strirng

var ch [NJ]chan carta  // NJ canais de itens tipo carta  

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta, ... ) {
	mao := cartasIniciais    // estado local - as cartas na mao do jogador
	nroDeCartas := M         // quantas cartas ele tem 
    cartaRecebida := " "     // carta recebida é vazia

	for {
		if // tenho que jogar
			{
			fmt.Println(id, " joga") // escreve seu identificador
			// escolhe alguma carta para passar adiante ...
			// guarda carta que entrou 
			// manda carta escolhida o proximo
			out <- cartaParaSair        
		} else {  
			// ...
			cartaRecebida := <-in   // recebe carta na entrada
			//  ...
			// e se alguem bate ?
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inicialize a semente para a geração de números aleatórios

	cartasEscolhidas := []carta{"A", "B", "C", "D"} // cartas iniciais de cada jogador
	for i := 0; i < NJ; i++ {
		ch[i] = make(chan struct{})
	}
	// cria um baralho com NJ*M cartas
	for i := 0; i < NJ; i++ {
		
		go jogador(i, ch[i], ch[(i+1)%N], cartasEscolhidas , ...) // cria processos conectados circularmente
	}

	for i := 0; i < NJ; i++ {
		// escolhe aleatoriamente (tira) cartas do baralho, passa cartas para jogador
        cartasEscolhidas := make([]carta, 4)
        for j := 0; j < 4; j++ {
            indiceAleatorio := rand.Intn(4) // Gera aleatoriamente um índice entre 0 e 3
            cartasEscolhidas[j] = carta('A' + indiceAleatorio)
        }

        go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas)
    }
	
	<-make(chan struct{}) // bloqueia
}
