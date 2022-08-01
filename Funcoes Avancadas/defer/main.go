package main

import "fmt"

func funcao1() {
	println("Executando a função 1")
}

func funcao2() {
	println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Verificando aprovação do aluno")
	fmt.Println("Aprovação de alunos!")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false

}

func main() {

	/* fmt.Println("Defer, adia a execussão de um pedaço de código. Adicionando a nomeclatura a frente do nome da função, ele a faz executar posteriormente, ignorando a sua ordem de chamada.")
	 */
	/* 	defer funcao1()
	   	// defer = Adiar
	   	funcao2() */

	fmt.Println(alunoAprovado(8, 9))

}
