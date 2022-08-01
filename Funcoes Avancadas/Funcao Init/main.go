package main

import "fmt"

func main() {
	// Utilizada normalmente para setar um valor que pode ser acessado antes da variavel main, iniciar.

	fmt.Println("Função pae!")

}
func init() {
	fmt.Println("Função executada antes da função main.")
}
