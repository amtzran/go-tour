package main

import "fmt"

/**
Constantes numéricas
Las constantes numéricas son valores de alta precisión.
Una constante sin tipo toma el tipo que necesita su contexto.

En el ejemplo de la derecha intente imprimir needInt(Big) también.
(Un int puede almacenar como máximo un número entero de 64 bits y, a veces, menos).
*/

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
