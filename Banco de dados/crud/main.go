package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ThailanTec/banco-de-dados/crud/servidor"
	"github.com/gorilla/mux"
)

type cota struct {
	Name string `json:"nome"`
	Age  string `json:"age"`
	Skin string `json:"skin"`
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/usuarios", servidor.CriarUser).Methods(http.MethodPost).Methods(http.MethodPost)

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))

}
