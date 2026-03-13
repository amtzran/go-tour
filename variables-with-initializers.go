package main

import "fmt"

/**
Variables con inicializadores
Una declaración var puede incluir inicializadores, uno por variable.
Si hay un inicializador, se puede omitir el tipo; la variable tomará el tipo del inicializador.
Como se muestra en el ejemplo de la derecha se inicializan sin tipo.
*/

var i, j int = 1, 2
var a = 10

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(a, i, j, c, python, java)
}
