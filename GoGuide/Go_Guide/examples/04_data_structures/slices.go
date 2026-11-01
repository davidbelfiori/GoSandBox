// slices.go - Slice come vista su un array sottostante (slide 29-30).
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// s[i:j] include l'elemento i ed esclude j.
	var s []int = primes[1:4]
	fmt.Println(s, "len =", len(s), "cap =", cap(s)) // [3 5 7] len=3 cap=5

	// La slice NON copia i dati: modificare s modifica l'array.
	s[0] = 100
	fmt.Println("array dopo s[0]=100:", primes)

	// Due slice sulla stessa porzione si influenzano a vicenda.
	t := primes[:2]
	fmt.Println("t =", t) // [2 100]

	// La capacità arriva fino alla fine dell'array sottostante.
	u := s[:5]
	fmt.Println(u, len(u), cap(u))

	// Controllo dei limiti a runtime: fuori dalla capacità -> panic.
	defer func() {
		fmt.Println("recuperato:", recover())
	}()
	_ = s[:6] // slice bounds out of range
}
