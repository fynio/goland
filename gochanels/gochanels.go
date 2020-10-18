package main

import "fmt"

func main() {

	// este canal esta definido como string
	channel := make(chan string)

	go func(channel chan string) {

		// simulamos un ciclo infinito
		for {
			var name string

			// aqui espera
			fmt.Scanln(&name)

			// aqui la enviamos
			channel <- name
		}
	}(channel)

	for {
		// ya rutina se detiene por que esta esperando algo que le ASIGNE el canal
		msg := <-channel

		// continua esto
		fmt.Println("Estoy imprimiento lo que recibi del canal: " + msg)

	}

}
