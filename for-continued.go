package main

import "fmt"

/**
For continuo
La instrucción inicial y posterior son opcionales.
*/

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
