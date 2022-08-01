package main

import (
	"fmt"
)

/* A função vai retornar um numero que está em determinada posição  da sequencia efibonatica(onde o proximo numero é sempre a soma dos outros 2 numeros anteriores.).
Deve-se determinar onde essa função deve parar de executar, para não causar estouro de pilha (o sistema não vai comportar), não recomendado para numeros muito grande*/
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao

	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas, são funções que chamam elas mesmas, depende para que ela funcine, de uma outra execussão dela mesma.")

	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println("________________________")
	fmt.Println(fibonacci(posicao))

}
