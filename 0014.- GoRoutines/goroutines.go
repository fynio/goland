package main

import (
	"fmt"
	"strings"
	"time"
)

/*
func main() {
	minombrelentooo("Rodrigo")
	fmt.Println("queee aburrridooo")
}
*/

func main() {
	go minombrelentooo("Rodrigo")
	fmt.Println("queee aburrridooo")
	var wait string
	fmt.Scanln(&wait)
}

func minombrelentooo(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}
