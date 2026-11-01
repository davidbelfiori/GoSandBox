// named_returns.go - Valori di ritorno con nome e "naked return" (slide 16).
package main

import "fmt"

// x e y sono variabili locali dichiarate nella firma, inizializzate al
// valore zero. Il "return" nudo restituisce i loro valori correnti.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // equivale a: return x, y
}

func main() {
	fmt.Println(split(17))
}
