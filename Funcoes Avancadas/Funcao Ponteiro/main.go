package main

import "fmt"

/* Quando passamos um valor, sem fazer o apontamento com um ponteiro, estamos passando uma copia desse mesmo. Logo, não estamos alterando a variavel.
Para inverter o valor devemos apontar atraves de um ponteiro de int(*int), então vamos alterar o endereço de memoria.
*/

func polos(num *int) {
	// Quando passamos um ponteiro, estamos passando uma referencia para a função e não uma copia.
	*num = *num * -1

}

func main() {

	number := 20
	fmt.Println(number)

	polos(&number)

	println(number)
}
