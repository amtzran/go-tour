package main

import "fmt"

/**
Funciones
Una función puede tomar cero o más argumentos.

En el ejemplo que se muestra a la derecha, add toma dos parámetros de tipo int.
Observar que el tipo viene después del nombre de la variable.
*/

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
