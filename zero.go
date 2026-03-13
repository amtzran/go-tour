package main

import "fmt"

/**
Valores cero

Las variables declaradas sin un valor inicial explícito reciben
su valor cero (zero value) por defecto.
El valor cero es:
	- 0 para tipos numéricos
	- false para tipo booleano
	- "" (la cadena vacía) para cadenas de texto
*/

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
