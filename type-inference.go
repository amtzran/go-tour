package main

import "fmt"

/**
Inferencia de tipo
Al declarar una variable sin especificar un tipo explícito (ya sea utilizando la sintaxis de expresión := o var =),
el tipo de variable se deduce del valor del lado derecho.

Cuando se escribe el lado derecho de la declaración, la nueva variable es del mismo tipo:

var i int
j := i // j es int
Pero cuando el lado derecho contiene una constante numérica sin tipo, la nueva variable puede
ser int, float64 o complex128 dependiendo de la precisión de la constante:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
En el ejemplo de la derecha intente cambiar el valor inicial de v en el código de ejemplo
y observe cómo se ve afectado su tipo.
*/

func main() {
	v := 42 // cambiame!!
	fmt.Printf("v is of type %T\n", v)
}
