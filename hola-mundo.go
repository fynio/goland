package main

import "fmt"

func holamundo(s string) string {
	return "Hola " + s
}

func main() {
	fmt.Println(holamundo("Mundo"))
}
