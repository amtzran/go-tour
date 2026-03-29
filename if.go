package main

import (
	"fmt"
	"math"
)

/**
If
La declaración if de Go son como si fueran bucles for; la expresión no necesita
estar rodeada de paréntesis ( ) pero las llaves { } son requeridas.
*/

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
