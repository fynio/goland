package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	executeReadFile()

	fmt.Println("nunca me voy a imprimir")
}

func executeReadFile() {
	// con ejecucion validamos que se leea el archivo
	ejecucion := readfile()

	fmt.Println("Se ejecuto correctamente", ejecucion)
}

func readfile() bool {

	archivo, error := os.Open("./archivco.txt")

	//DEFER se ejecuta sin importar el punto donde se ejecute el return
	defer func() {
		archivo.Close()
		fmt.Println("defer si se ejecuta")

		r := recover()
		fmt.Println(r)

	}()

	if error != nil {
		fmt.Println("hay un error")

		//PANIC  imprimer un error a detalle
		panic(error)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		time.Sleep(100 * time.Millisecond)
		fmt.Println(linea)
	}

	if true {
		return true
	}

	fmt.Println("nunca me ejecuto")
	archivo.Close()

	return true

}
