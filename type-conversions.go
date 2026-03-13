package main

import (
	"fmt"
	"math"
)

/**
Conversiones de tipo
La expresión T(v) convierte el valor v al tipo T.
Algunas conversiones numéricas:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
O, dicho más simple:

i := 42
f := float64(i)
u := uint(f)
A diferencia de C, en Go la asignación entre elementos de diferente tipo requiere una conversión explícita.
*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
