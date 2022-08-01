package main

import "fmt"

func main() {

	// Maps chave e valor mais rigidos, chave tem sempre o mesmo tive e não é uma estrutura de valor mutavel(alteravel).

	fmt.Println("Maps")

	// Criar maps dentro do colchetes é o tipo da chave e fora do colchete é o tipo dos valores.

	/* user := map[string]string{
		"nome":      "Thailan",
		"sobrenome": "Santos",
	}

	fmt.Println(user) */

	// Alinhar maps, pode passar varias estruturas de maps, para conseguir de uma forma mais detalhada passar informações

	user3 := map[string]map[string]string{

		"dados": {
			"Primeiro Nome": "Thailan",
			"Ultimo Nome":   "Jesus",
		},
		"graduacao": {
			"Faculdade": "Unopar",
			"Curso":     "Analise e Desenvolvimento de sistemas",
		},
	}

	fmt.Println(user3)
	// Função basica da linguagem Go, onde temos como deletar um tipo de item.
	delete(user3, "graduacao")
	fmt.Println(user3)
}
