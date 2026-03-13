package main

import (
	"fmt"
	"math"
)

/**
Nombres exportados
En Go, un nombre se exporta si comienza con una letra mayúscula.
Por ejemplo, Pizza es un nombre exportado, al igual que Pi, que se exporta desde el paquete math.

Aquí pizza y pi no comienzan con una letra mayúscula, por lo que no se exportan.

Al importar un paquete, solo se puede hacer referencia a sus nombres exportados.
Cualquier nombre "no exportado" no es accesible desde fuera del paquete.
*/

func main() {
	fmt.Println(math.Pi)
}
