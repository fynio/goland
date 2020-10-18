package main

import "fmt"

type human struct {
	name string
}

func (humano human) hablar() string {
	return "bla, bla, bla"
}

type auxiliar struct {
	name string
}
type tutor struct {
	human
	auxiliar
}

// sobreescribimos los metodos
func (tutor_ tutor) hablar() string {
	return tutor_.human.hablar() + " Hola mundo"
}
func main() {
	_tutor := tutor{human{"Rodrigo"}, auxiliar{"Fynio"}}
	fmt.Println(_tutor.human.name)
	fmt.Println(_tutor.hablar())
}
