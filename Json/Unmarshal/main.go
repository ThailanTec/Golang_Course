package main

import (
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

	// Se utilizar ao invez do json:"nome", podemos usar json"-", q ele não vai ser convertido.
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Slink","raca":"Basset","idade":1}`

	var c Cachorro

	// A função Unmashal ela recebe 2 paremetros 1- os dados e 2 - o endereço de memoria(&), por quer desejamos que ela fique alterada. Ele não recebe o parametro de forma correta, que é slice de bytes. Para receber um slice de byyes, podemos fazer bytes[cachorro]  e assim passamos para ele um slice de bytes. E devemos passar também o erro, e trata-lo em seguida.
	err := json.Unmarshal([]byte(cachorroEmJson), &c)

	if err != nil {
		log.Fatal("Erro pra krlh")
	}

	fmt.Println(c)

	// Para converter para um map, pode-se fazer da seguinte forma: obs: Tem que tomar cuidado, pois tem que verificar o tipo que está vindo.
	cachorroEmJson2 := `{"nome":"Luck","raca":"Basset","idade":"1"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJson2), &c2); erro != nil {
		log.Fatal("Deu erro mano")
	}
	fmt.Println(c2)
}
