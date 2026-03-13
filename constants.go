package main

import "fmt"

/**
Constantes
Las constantes se declaran como variables, pero con la palabra clave const.
Las constantes pueden ser caracteres, cadenas de texto, booleanos o valores numéricos.

Las constantes no se pueden declarar usando la sintaxis :=.
*/

const Pi = 3.14

func main() {
	const World = "México 🇲🇽"
	fmt.Println("Hello", World)
	fmt.Println("Happy day of", Pi)

	const Truth = true
	fmt.Println("Go manda?", Truth)
}
