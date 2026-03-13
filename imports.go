package main

import (
	"fmt"
	"math"
)

/**
Importaciones
El código que se muestra a la derecha agrupa las importaciones
en una instrucción de importación "factorizada" entre paréntesis.

También se puede escribir varias declaraciones de importación, como:
import "fmt"
import "math"
Pero es un buen estilo o práctica usar la declaración de importación
factorizada cuando son multiples importaciones.
*/

func main() {
	fmt.Printf("Ahora tienes %g problemas.\n", math.Sqrt(7))
}
