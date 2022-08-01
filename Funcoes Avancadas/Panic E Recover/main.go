package main

import "fmt"

func recupera() {

	/* Utilizamos a variavel R para recuperar o programa, logo em seguida passamos para ele o valor de recover e se o mesmo for diferente de nulo a aplicação deve continuar. Significando que a recuperação do programa deu certo. */
	if r := recover(); r != nil {
		fmt.Println("Recuperar execucão")
	}

}

func alunoApro(n1, n2 int) bool {

	defer recupera()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	/*Quando utilizado o programa para completamente com a sua execussão. Não tendo como realizar uma tratativa, diferente do erro que tem como tratar.*/
	panic("Fodeu mano")
	/*A tag recover, serve para que possamos pegar novamente o programa enquanto está em panico, dessa forma a aplicação continua correndo.*/
}

func main() {
	fmt.Println("Interronpe a funcionalidade do programa e para TUDO.")

	fmt.Println(alunoApro(6, 6))

	fmt.Println("Pós execussão!")
}
