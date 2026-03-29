package main

import (
	"fmt"
	"math"
)

/*
If con una breve declaración
Al igual que for, la declaración if puede comenzar con una breve declaración
para ejecutar antes de la condición.
Las variables declaradas por la instrucción solo están en el alcance hasta el final del if.
*/

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
