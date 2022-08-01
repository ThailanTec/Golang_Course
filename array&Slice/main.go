package main

import (
	"fmt"
)

func main() {

	fmt.Println("Slice E Array")

	// Lista de valores / Array, pode criar das 2 formas

	/* 	var array1 [5]int

	   	array1[0] = 11
	   	array1[1] = 12
	   	fmt.Println(array1)

	   	array2 := [2]string{"Thailan", " Thamos"}

	   	fmt.Println(array2)

	   	// Removendo a restrição do array, utliza-se os 3 pontos

	   	array3 := [...]int{2, 2, 4, 5, 5, 45, 5, 55}

	   	fmt.Println(array3)
	*/
	// Slices

	slice := []string{"Cs:go", "Valorant", "PB"}

	// Adicionando itens 'Append' = ela adiciona o item no slice e retorna o slice com o item add

	slice = append(slice, "LOL", "Dota")
	fmt.Println(slice)

	// Arrays internos:

	fmt.Println("------------")

	/* Essa função Make, cria um slice ou outro item, e com base nisso podemos realizar testes ou outras coisas.
	Então, temos baixo um arry, com 15 posições e um slice com 10. Caso, houvesse um estouro da memoria passada, o go, dobra o valor para que tenhamos espaço de memoria.
	*/
	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
}
