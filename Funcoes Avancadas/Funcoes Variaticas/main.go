package main

import "fmt"

/*  Passando dessa forma( exemplo primeira função), estamos dizendo que ele pode receber multiplos parametros, não dizendo exatamente quantos irão ser. Mas o mesmo tem que ser sempre o ultimo e não pode ter outro além dele. */

func soma(numeros ...int) int {
	total := 0

	for _, numenumero := range numeros {
		total += numenumero
	}

	return total
}

func escreve(texto string, numero ...int) {
	for _, numeros := range numero {
		fmt.Println(texto, numeros)
	}
}

func main() {
	fmt.Println("Funções Variaticas: Funções na qual não precisa dizer quantos parametros ela vai receber.")

	totalDaSoma := soma(12, 74)

	fmt.Println(totalDaSoma)

	escreve("Olá mundo", 22)

}
