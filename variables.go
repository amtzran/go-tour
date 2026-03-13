package main

import "fmt"

/**
Variables
La instrucción var declara una lista de variables; como en las listas de argumentos
de función, el tipo es el último en la declaración.
Una sentencia var puede estar a nivel de paquete o de función. Ambos se muestran en el ejemplo de la derecha.
*/

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
