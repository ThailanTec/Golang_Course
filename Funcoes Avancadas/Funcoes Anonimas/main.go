package main

import "fmt"

func main() {
	fmt.Println("Funções que não tem nome")
	/*Para executar uma função do tipo anonima podemos fazer da forma abaixo, colocar no final da mesma uma chave, onde o go vai reconhecer e inciar a mesma
	Podemos também passar parametros.*/

	func() {
		fmt.Println("Sem Parametro!")
	}()

	func(texto string) {
		fmt.Println(texto)

	}("Olha o Parametro dela aquiooo!")
}
