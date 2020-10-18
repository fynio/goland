package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileData, err := ioutil.ReadFile("./archivo.txt")

	if err != nil {
		fmt.Println("hubo un error")
	}

	fmt.Println(string(fileData))
}
