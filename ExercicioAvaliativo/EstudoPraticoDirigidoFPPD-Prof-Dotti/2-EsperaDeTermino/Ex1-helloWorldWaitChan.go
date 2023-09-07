// Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// >>> Veja o Ex0 desta série
// ABRE E FECHA CONCORRENCIA
// Há várias formas de esperar o término de processos concorrentes.
// EXERCICIOS:
//   1)  isto seria uma solução para sincronizar o final do programa? 
//		Não pois continuaria saindo 10 e depois os outros 10. Mas garante que realmente saírão todas as palavras.
//   2)  aumente para criar 10 prodessos concorrentes say(...).
//       como voce faz a espera de todos ?
//		ele espera pq está dentro de um canal.
// OBS:  tente um comando de repeticao.

package main

import (
	"fmt"
)

func say(s string, c chan struct{}) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	c <- struct{}{}
}

func main() {
	fin := make(chan struct{})
	go say("world", fin)
	go say("hello", fin)
	<-fin
	<-fin
}
