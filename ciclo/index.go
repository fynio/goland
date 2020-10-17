package main

import "fmt"

func main() {

	/*
		for i := 0; i <= 10; i++ {

			fmt.Println("este es el número ", i)
		}
	*/

	/*
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
	*/

	/*
		i := 0
		for {
			fmt.Println(i)
			i++
			if i > 10 {
				break
			}
		}

	*/

	i := 0
	for {

		if i == 2 {
			i++
			continue
		}

		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}

}
