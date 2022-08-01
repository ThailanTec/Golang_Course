package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Intefaces com tipos Genericos! Você não precisa passar o tipo que a interface vai receber, mas pode causar problemas. Utilzar apenas em casos especificos.")

	generica("String")
	generica(1)
	generica(true)
}
