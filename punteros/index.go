package main

import "fmt"

func main() {

	/*
		1. Es una direccion de memoria
		2. En lugar del valor, tenemos la direccion en la que esta el valor
		3. X,Y => AS123D =>5
		4. X )> AS123D => 6
		5. y ¿? => 6

		// t es igual al tipo de dato
		*T => *int, *stringm *Struct

	*/

	var x, y *int
	entero := 5

	// & accede a la direccion de memoria no al valor
	x = &entero
	y = &entero

	// estamos alterando la direccion de memoria
	*x = 6

	// por lo tanto cuando alteramos la direcciond de memoria y x y y  tienen la misma
	// pues entonces los dos tienen el mismo valor

	fmt.Println("La direccion de memoria de x es :", x)
	fmt.Println("La direccion de memoria de y es :", y)

	fmt.Println("La direccion de memoria de x es :", *x)
	fmt.Println("La direccion de memoria de y es :", *y)

}
