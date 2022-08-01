package main

import (
	"fmt"
	"time"
)

func main() {

	/* Concorrencia vs Paralelismo

	Paralelismo: Quando se tem duas ou mais tarefas sendo executadas oa mesmo tempo. Só sendo possivel , quando se tem um processador com mais de um nucleo, distribuindo a tarefa entre os nucleos.

	Concorrencia: Não necessariamente são executadas ao mesmo tempo, porem podem está dividida entre os processadores multinuclear ou no mesmo nucleo. Fazendo um revezamento no nucleo entre uma e outra.

	*/

	/*No exemplo abaixo, as funções vão ficar sendo executadas de forma infinito, mas não pode ficar dessa forma, logo vamos usar goroutines,
	para dividir o processamento das tarefas.
	Para chamar a go routine, chamamos o 'go', propiedade que informa para o programa que ele deve seguir com a sua execussão, mesmo, antes que a atividade tenha finalizado. A ordem pode mudar, pois estão rodando ao mesmo tempo.

	Go Routines: Funções metodos que podem ser aplicados para que aplicação continue antes que uma chamada seja finalizada.

	*/

	go escrever("Olá mundo!")
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
