package main

import "fmt"

/* Vamos criar um padrão de código para usar gorotines */

/* Utilizando o worker, conseguimos chamar os canais mais de uma vez, dessa forma trazendo mais velocidade na execução de uma função recursiva. */
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao

	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	tarefas := make(chan int, 45)
	resultado := make(chan int, 45)

	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultado

		fmt.Println(resultado)
	}

}

/* Para Informar que o canal envia ou recebe dados, é passado as seguintes informações: tarefas <-chan, dessa forma, dizemos que o canal vai receber dados e para dizer que ele envia dados passados a seta depois resultado chan<-. Se não informar nada, o canal é declarado como bi direcional*/
func worker(tarefas <-chan int, resultado chan<- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}
