package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de repetiação for/ Maneiras de usar o for")

	// 1 Padrão - Enquanto for verdadeiro, continua rodando, quando bater, para de rodar auto.
	/*
		for i := 0; i < 10; i++ {

			time.Sleep(time.Second)
			fmt.Println("Verdade")
		} */
	// 2 Range - Utliza-se quando vai inteirar em algo, array, slice, map, usado nesses cenarios

	nomes := [3]string{"João", "Douglas", "Tulipio"}

	/* Retorrna o indice (posição) e o valor(dados), por padrão o primeiro
	valor é sempre o indice. Podemos utilizar o "_", para ignorar o primeiro valor. */

	// Com slice/Array
	for _, nome := range nomes {

		fmt.Println(nome)

	}

	// Com maps

	user := map[string]string{
		"nome":  "Jusepio",
		"idade": "22",
	}
	fmt.Println(user)

	for _, valor := range user {
		fmt.Println(valor)
	}

	// For infinito kkkk

	for {
		fmt.Println("Ao infinito e além!")
	}
	// OBS: NÃO DÁ PARA FAZER COM STRUCT POW!

}
