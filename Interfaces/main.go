package main

import (
	"fmt"
	"math"
)

/* Muito utilizada em Go, quando precisamos de flexibilidade com os tipos. */
// Criação de uma interface

/* Ela só tem assinaturas de metodos, logo, você não pode preencher ela da mesma forma que faria o preencheimento de um struct*/

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {

	fmt.Printf("A área da forma é %0.2f \n", f.area())

}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main2() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
