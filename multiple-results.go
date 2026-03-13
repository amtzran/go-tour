package main

import "fmt"

/**
Múltiples resultados
Una función puede devolver cualquier número de resultados.
La función swap mostrada a la derecha devuelve dos cadenas de texto.
*/

func swap(x, y string) (string, string) {
	return y, x
}

func calculateTotal(subtotal, discount float64) (float64, string) {
	total := subtotal + discount
	if subtotal < 0 {
		return total, "El subtotal no puede ser 0."
	}
	return total, "Total calculado."
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	total, message := calculateTotal(100, 30)
	fmt.Println(total, message)
}
