package main

import "fmt"

type estudante struct {
	name   string
	idade  int
	altura float32
}

// conseguimos herdar o tipo da struct de cima e reutilizar, tipo uma henrança familiar de primos kkkkk
type pessoa struct {
	estudante
	curso string
}

func main() {
	println("Olá")

	aluno1 := estudante{"Thailan", 21, 1.75}

	pessoa2 := pessoa{aluno1, "Engenharia"}

	fmt.Println(aluno1, pessoa2)

}
