// slice_pitfalls.go - Errori comuni con array e slice.
package main

import "fmt"

// Errato: append può restituire una slice con un nuovo array; il
// risultato va riassegnato. Qui la modifica si perde per il chiamante,
// perché s è una copia dell'header (puntatore, len, cap).
func addWrong(s []int) {
	s = append(s, 99)
}

// Corretto: restituire la slice aggiornata.
func addRight(s []int) []int {
	return append(s, 99)
}

func main() {
	s := []int{1, 2, 3}
	addWrong(s)
	fmt.Println("dopo addWrong:", s) // [1 2 3]
	s = addRight(s)
	fmt.Println("dopo addRight:", s) // [1 2 3 99]

	// append senza riassegnare: il compilatore stesso lo rifiuta
	// ("append(s, 4) (value of type []int) is not used"). Va scritto
	// s = append(s, 4).

	// Una funzione può però modificare gli ELEMENTI esistenti, perché la
	// slice condivide l'array sottostante.
	double(s)
	fmt.Println("dopo double:", s)

	// Array e slice sono diversi: l'array si copia per valore.
	arr := [3]int{1, 2, 3}
	arrCopy := arr
	arrCopy[0] = 100
	fmt.Println(arr, arrCopy)
}

func double(s []int) {
	for i := range s {
		s[i] *= 2
	}
}
