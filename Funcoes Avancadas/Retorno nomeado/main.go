package main

import "fmt"

func calcuM(n1, n2 int) (soma int, sub int) {

	soma = n1 + n2
	sub = n2 - n1

	return
}

func main() {

	fmt.Println("Funções com retorno nomeado")
	soma, sub := calcuM(10, 20)

	fmt.Println(soma, sub)

}
