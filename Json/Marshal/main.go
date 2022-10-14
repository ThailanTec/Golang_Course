package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

/*
		MARSHAL:
	Podemos usar para converter um map para json ou uma struct para json
*/

/* UNMARSHAL:
Usamos para transformar um json em um struct ou map dependendo da necessidade
*/

type Cachorro struct {
	/* Usamos para dizer que o json vai receber o nome lá na saida do json: `json:"nome"`, assim vai ficar essa nomeclatura e não o que declaramos. */
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	c := Cachorro{"Slink", "Basset", 1}

	// Para realizar o Marshal ( passar para json), podemos fazer da seguinte forma:

	slinkemJson, err := json.Marshal(c)

	// Mas para vermos a saida correta, precisamos usar o pacote da do go, para que eles nos mostra da forma correta o json. Bytes, que é o bytes.NewBuffer, que vai receber um array de bytes e logo em seguida vai passar a saida correta: fmt.Println(slinkemJson)

	if err != nil {
		log.Fatal("DEU ERRO AQUI, NÂO VIROU JSON NÃO!")
	}

	fmt.Println(bytes.NewBuffer(slinkemJson))

	c2 := map[string]string{
		"nome":  "Luck",
		"idade": "1.5",
	}
	luckemJson, err := json.Marshal(c2)
	if err != nil {
		log.Fatal("DEU ERRO AQUI, NÂO VIROU JSON NÃO!")
	}
	fmt.Println(bytes.NewBuffer(luckemJson))

}
