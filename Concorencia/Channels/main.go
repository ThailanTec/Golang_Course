package main

import (
	"fmt"
	"time"
)

/* Canais/Channels: São chamados assim por que é um canal de comunicação. Pode enviar ou receber dados, e sincronizar as funções entre as rotinas. */

func main() {

	/* Para criar um canal utilizamos a função make, na qual fica responsavel por criar. Logo em seguida passamos a chave que no caso do canal é: 'chan' e logo em seguida o tipo de dados que vai trafegar entre eles. */

	canal := make(chan string)

	go escrever("Olá meu mundo", canal)

	for msg := range canal {
		fmt.Println(msg)
	}
	fmt.Println("Fim do programa.")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)

}
