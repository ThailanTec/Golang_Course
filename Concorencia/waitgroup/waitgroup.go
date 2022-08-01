package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/* Utilizamos go routines para não esperar uma terminar para a outra iniciar, assim podemos sincronizar as mesmas finalizar, antes do programa terminar de executar.

	1 - Criamos uma variavel e passamos como waitGroup e logo em seguida passamos a função padrão padrão do go, 'sync', que vai ser responsavel por sincronizar as coisa.

	2 - Depois passamos o waitGroup.Add(), e informamos quantas rotinas temos na nossa aplicação que vão trabalhar juntas.

	3 - Logo em seguida criamos funções anonimas e dentro dela passamos o código que desejamos que entre dentro das rotinas.

	4- Seguindo ainda dentro da função waitGroup.Done(), no final do código, informando que o mesmo já acabou a sua execussão.

	5 - Passamos o waitGroup.Wait(), para falar a função main( ou a função onde estiver dentro), para esperar a execucão chegar em 0, antes de seguir.

	*/
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
