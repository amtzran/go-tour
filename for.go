package main

import "fmt"

/**
For
Go solo tiene una estructura de bucle, el bucle for.

El bucle for básico tiene tres componentes separados por punto y coma:

La declaración inicial: ejecutada antes de la primera iteración
La expresión de la condición: evaluada antes de cada iteración
La declaración posterior: ejecutada al final de cada iteración
La declaración de inicio será a menudo una declaración de variable corta, y las variables
declaradas allí son visibles solo en el ámbito de la declaración for.

El loop se detendrá una vez que la condición booleana evalué a falso.

Nota: a diferencia de otros lenguajes como C, Java, o JavaScript no hay paréntesis alrededor de
los tres componentes de la declaración for y las llaves { } son siempre requeridas.
*/

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
