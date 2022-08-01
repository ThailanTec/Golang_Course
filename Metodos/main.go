package main

import "fmt"

/* Diferente de funções metodos, estão obrigatoriamente ligados a alguma coisa, como uma struct, interface, etc  */

type user struct {
	Nome  string
	Idade uint
}

/* Assim, então escrevemos como vai ser a estrutura de um metodo. Estamos dizendo que todos os usuarios tem um metodo com o nome: salvar.
A letra 'u' é uma variavel que podemos utilizar para acessar esse metodo.
*/
func (u user) salvar() {
	fmt.Printf("Salvando os dados do %s no banco de dados.\n", u.Nome)
}

/* Quando precisamos alterar algum dado na struct, podemos utilizar metodos como a forma baixo, para definir qual vai ser o melhor caminho para o mesmo.
Via de regra quando você tem um metodo que vai fazer alguma alteração no seu struct, você costuma passar a referencia ( ponteiro*), para o seu metodo.
*/

func (u *user) niver() {

	u.Idade++

}

func main() {

	usuario := user{"Thailan", 22}
	fmt.Println(usuario)
	usuario.salvar()
	usuario.niver()
	fmt.Println(usuario)

}
