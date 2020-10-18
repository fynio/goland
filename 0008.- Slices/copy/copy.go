package main

import "fmt"

func main() {

	/*
		slice := []int{1, 2, 3, 4}
		copia := make([]int, 4)

		fmt.Println(slice)
		fmt.Println(copia)
	*/

	// LA FUNCION COPIA COMPIA EL MINIMO DE ELEMENTOS

	/*
		slice := []int{1, 2, 3, 4}
		copia := make([]int, 4)

		copy(copia, slice)
		fmt.Println(slice)
		fmt.Println(copia)
	*/

	/*
		slice := []int{1, 2, 3, 4}
		copia := make([]int, 0)

		copy(copia, slice)
		fmt.Println(slice)
		fmt.Println(copia)
	*/

	// PARA CORREGIR ESE ERROR

	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice))

	copy(copia, slice)
	fmt.Println(slice)
	fmt.Println(copia)

}
