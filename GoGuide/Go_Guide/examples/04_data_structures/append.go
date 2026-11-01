// append.go - append e crescita della capacità (slide 32-33).
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0) // funziona su slice nil
	printSlice(s)

	s = append(s, 1) // la slice cresce quando serve
	printSlice(s)

	s = append(s, 2, 3, 4) // più elementi in una volta
	printSlice(s)

	// Capacità piena: viene allocato un nuovo array e gli elementi sono
	// copiati.
	t := make([]int, 3, 3)
	t = append(t, 4)
	fmt.Println(t) // [0 0 0 4]
	printSlice(t)

	// Capacità disponibile: nessuna nuova allocazione, e l'array
	// sottostante è condiviso.
	base := make([]int, 3, 10)
	x := append(base, 1)
	y := append(base, 2)    // sovrascrive la stessa posizione di x!
	fmt.Println(x[3], y[3]) // 2 2
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
