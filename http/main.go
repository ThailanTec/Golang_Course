package main

import (
	"log"
	"net/http"
)

func main() {

	/* Protocolo de comunicação - Base da comunicação WEB.

	Sendo Cliente > Servidor
	Tudo sendo

	Resquest -> Response

	Resquisição & Resposta

	Rotas: Forma de identificar o tipo de mensagem e o processamento que o servidor vai fazer em cima da mensagem.
	URI - Identificador de Recursos
	Método: Get, post, put, delete ( Buscar, Atualizar, atualizar, deletar dados.)

	*/

	// Utilizamos o handlerFunc para passar o URI o tipo de metodo que vai receber, assim quando acessarmos a url, ele vai escrever na tela essa mensagem na tela.

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo!"))
	})

	// Utilizamos para levantar um servidor de forma nativa do go
	log.Fatal(http.ListenAndServe(":5000", nil))

}
