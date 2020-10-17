package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22

	// Atoi CONVIERTE UN STRING A UN ENTERO
	edadStr := strconv.Itoa(edad)

	fmt.Println("Mi edad es " + edadStr)
}
