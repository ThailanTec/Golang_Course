package main

import "fmt"

/* São funções que referenciam variaveis que estão fora do seu corpo. Logo, se houver uma outra variavel com o mesmo nome, em sua 'memoria', ela não vai deixar que ocorra a alteração de tipo.*/

func clousure() func() {
	texto := "Dentro do bglh!"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("Funções do tipo Clousure.")
	texto := "Dentro do main mano!"

	fmt.Println(texto)

	funcaoNova := clousure()

	funcaoNova()
}
