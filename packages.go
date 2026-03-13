package main

import (
	"fmt"
	"math/rand"
)

/**
Paquetes
Cada programa de Go se compone de paquetes.

Los programas en Go comienzan a ejecutarse en el paquete main que siempre
debe existir como punto de entrada para el flujo del programa.

El programa que se muestra a la derecha está usando los paquetes con rutas de importación "fmt" y "math/rand".

Por convención, el nombre del paquete es el mismo que el último elemento de la ruta de importación.
Por ejemplo, el paquete "math/rand" comprende archivos que comienzan con la declaración package rand.
En cambio, "math" viene siendo parte de la ruta donde se encuentra el paquete y a su vez también
puede ser un paquete que comprende archivos que comienzan con la declaración package math.
*/

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
