package main

import (
	"fmt"
	"math/rand"
)

/**
For Infinito
Si omite la condición del bucle, se repetirá para siempre, por lo que un bucle infinito
se expresa de forma compacta.
*/

func main() {
	for {
		randomNumber := rand.Intn(100)
		if randomNumber%2 == 0 {
			fmt.Println("Even:", randomNumber)
			break
		}
	}
}
