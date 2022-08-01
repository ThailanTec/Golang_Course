package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida!")

	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args) // utilizado para a maquina aceitar os comandos.
	if erro != nil {

		log.Fatal(erro)

	}
}
