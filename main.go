// Criar struct com todos os tipos do go!

package main

import (
	"fmt"
	"time"
)

type formas struct {
	Nome      string
	Idade     int32
	Casa      int64
	Cep       uint
	Salario   float32
	Desconto  float64
	feliz     bool
	CreatedAt time.Time
}

type forma2 struct {
	Nome      string    `json:"name2"`
	Idade     int32     `json:"idade"`
	Casa      int64     `json:"casa"`
	Cep       uint      `json:"cep"`
	Salario   float32   `json:"salario"`
	Desconto  float64   `json:"desconto"`
	Feliz     bool      `json:"feliz"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {

	pessoa := formas{"Thailan", 22, 887, 32424040, 22.500, 533.50, true, time.Now()}

	fmt.Println(pessoa)

	pessoa2 := forma2{
		Nome:      "Jefferson",
		Idade:     38,
		Casa:      228,
		Cep:       827282882,
		Salario:   8100,
		Desconto:  2000,
		Feliz:     true,
		CreatedAt: time.Now(),
	}
	fmt.Println(pessoa2)
}
