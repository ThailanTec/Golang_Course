package main

import (
	"fmt"
	"time"
)

/* Poder pegar 2 ou mais canais e juntar em apenas um canal. */

func main() {

	canal := multiplexar(<-escrever("Olá mundo"), <-escrever("Olá mundo 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)

	}

}

func multiplexar(canalEnt1, canalEnt2 string) <-chan string {

	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEnt1:
				canalSaida <- mensagem
			case mensagem := <-canalEnt2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida

}

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {

		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}

	}()

	return canal
}
