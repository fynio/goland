package main

import "fmt"

func main() {
	// cuando no se le asigna el tamaño lo convierte automaticamente en un slice
	// var matriz []int

	/*
		matriz := []int{1, 2, 3, 4, 5}

		fmt.Println(matriz)
	*/

	arreglo := [3]int{1, 2, 3}

	// CON SLICE TAMBIEN PUEDES PARTIR EL ARRREGLO
	/*
		//OUTPUT:2  DEL ARRELGO  [1,2,3]
		slice := arreglo[1:2]
	*/

	/*
		//OUTPUT:1,2  DEL ARRELGO  [1,2,3]
		slice := arreglo[0:2]
	*/

	/*
		//OUTPUT:1,2,3  DEL ARRELGO  [1,2,3]

	*/
	slice := arreglo[0:3]
	fmt.Println(slice)
}
