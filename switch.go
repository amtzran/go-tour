package main

import (
	"fmt"
	"runtime"
)

/*
Switch
Una sentencia switch es una forma corta de escribir una secuencia de sentencias if - else.
Ejecuta el primer caso cuyo valor es igual a la expresión de condición.

El switch de Go es como el de C, C++, Java, JavaScript y PHP, excepto que Go solo ejecuta
el caso seleccionado, no todos los casos que le siguen.
En efecto, la sentencia break que se necesita al final de cada caso en esos lenguajes se
proporciona automáticamente en Go.
Otra diferencia importante es que los casos switch de Go no necesitan ser constantes y los
valores involucrados no necesitan ser enteros.
*/

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
