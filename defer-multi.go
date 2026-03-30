package main

import "fmt"

/*
Apilando defers
Las llamadas a funciones diferidas se introducen en una pila.
Cuando una función retorna, sus llamadas diferidas se ejecutan de último en entrar, primero en salir.
https://blog.golang.org/defer-panic-and-recover
*/

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
