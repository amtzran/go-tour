package main

import (
	"fmt"
)

/**
For es el while de Go.
En ese punto puedes soltar los puntos y coma: El while de C se escribe for en Go.
*/

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
