package main

import "fmt"

type user interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type admin struct {
	nombre string
}

func (administrador admin) Permisos() int {
	return 5
}

func (administrador admin) Nombre() string {
	return administrador.nombre
}

func auth(usuario user) string {
	if usuario.Permisos() > 4 {
		return usuario.Nombre() + "tiene permisos de administrador"
	}

	return ""
}

func main() {
	admin := admin{"rodrigo"}

	fmt.Println(auth(admin))
}
