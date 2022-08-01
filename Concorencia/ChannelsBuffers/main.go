package main

import "fmt"

func main() {

	/* Canais com Buffers, utilizado para poder especificar uma capacidade para o seu canal. ex: 2, diferente do canal normal, o canal com buffers, só vai bloquear quando o mesmo atingir a capacidade total dele. ObsP: Quando passamos o canal com 2, significa que o nosso canal tem 2 'portas' para receber de dados por exemplo, então, ele aceita até 2, passando disso, ele volta a dar o erro de deadlock.*/

	canal := make(chan string, 2)

	canal <- "Olá mundo!"

	msg := <-canal

	fmt.Println(msg)
}
