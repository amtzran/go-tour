package main

import "fmt"

/**
Declaración corta de variables
Sólo dentro de una función, la declaración de asignación corta := se puede usar
en lugar de una declaración var con tipo implícito.

Fuera de una función, cada declaración comienza con una palabra clave (var, func, etc.), por lo que la
construcción := no está disponible en ese ámbito superior.
*/

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
