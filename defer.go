package main

import "fmt"

/*
Defer
Una instrucción defer, aplaza la ejecución de una función hasta el retorno
de la función circundante (que la contiene).
Los argumentos de la llamada diferida se evalúan inmediatamente, pero la llamada a la
función no se ejecuta hasta que la función que la rodea retorna.
*/

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
