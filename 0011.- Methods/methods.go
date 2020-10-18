package main

import "fmt"

type user struct {
	edad             int
	nombre, apellido string
}

func (usuario user) nombreCompleto() string {
	return usuario.nombre + "	" + usuario.apellido
}

func (usuario *user) setName(n string) string {
	usuario.nombre = n

	return usuario.nombre
}

func main() {

	rodrigo := new(user)

	rodrigo.edad = 30
	rodrigo.nombre = "Rodrigo"
	rodrigo.apellido = "Garcia"

	rodrigo.setName("Fynio")
	fmt.Println(rodrigo.nombre)
}
