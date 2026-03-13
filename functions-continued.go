package main

import "fmt"

/**
Funciones continuación
Cuando dos o más parámetros de función con nombre consecutivos comparten
el mismo tipo, se puede omitir el tipo de todos menos del último.

En el ejemplo, se acorta
x int, y int
a:
x, y int
*/

func addSharedType(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(addSharedType(42, 13))
}
