package main

import "fmt"

/**
Valores de retorno con nombre
Los valores de retorno de Go pueden ser nombrados. Si es así, se tratan
como variables definidas en la parte superior de la función.

Estos nombres deben utilizarse para documentar el significado de los valores devueltos.
Una instrucción return sin argumentos devuelve los valores de retorno nombrados.
Esto se conoce como retorno "desnudo" ("naked" return).

Las declaraciones de retorno desnudas deben usarse solo en funciones cortas, como en el
ejemplo que se muestra aquí a la derecha. En funciones más largas pueden dañar la legibilidad.
*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(42))
}
