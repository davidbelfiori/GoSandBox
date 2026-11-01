// range.go - Iterare con range su una slice (slide 35).
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// Per array e slice, range produce indice e copia del valore.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
