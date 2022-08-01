package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Utilizamos o select (muito parecido com o switch), criamos um case, informando para quem receber primeiro, receber o dado primeiro é exibido primeiro, com definimos que o sistema vai receber o dado 1 primeiro, então, ele recebe primeiro e logo em seguida o canal 2 recebe. Assim temos uma maior velocidade no recebimento dos dados.
	*/
	canal1, canal2 := make(chan string), make(chan string)

	go func() {

		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {

		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	select {
	case msgCanal1 := <-canal1:
		fmt.Println(msgCanal1)
	case msgCanal2 := <-canal2:
		fmt.Println(msgCanal2)
	}

	for {
		msgCanal1 := <-canal1
		fmt.Println(msgCanal1)

		msgCanal2 := <-canal2
		fmt.Println(msgCanal2)
	}

}
