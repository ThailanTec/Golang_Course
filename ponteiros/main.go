package main

import "fmt"

func main() {
	fmt.Println("Salve")

	var variavel1 int = 10
	var variavel2 *int

	variavel2 = &variavel1

	println(*variavel2)
}
