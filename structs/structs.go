package main

import "fmt"

//type define su nuevo tipo
// tiene como nombre su segungo kword en este caso es user
// * deber ser en minusculas
type user struct {
	edad             int
	nombre, apellido string
}

func main() {

	rodrigo := user{edad: 20, nombre: "Rodrigo", apellido: "Garcia"}

	fmt.Println(rodrigo.nombre)

}
