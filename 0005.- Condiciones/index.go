package main

import "fmt"

func main() {
	x := 5
	y := 15

	if x > y {
		fmt.Printf("%d  es mayor que %d  \n", x, y)
	} else {
		fmt.Printf("%d  > %d \n", x, y)
	}
}
