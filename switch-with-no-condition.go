package main

import (
	"fmt"
	"time"
)

/*
Switch sin una condición es lo mismo que switch true.
Esta construcción puede ser una forma limpia de escribir largas cadenas if-then-else.
*/

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
