// structs.go - Struct (slide 28).
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// Letterale con campi nominati; i campi omessi valgono zero.
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v.X, v)

	// Un puntatore a struct: Go consente p.X al posto di (*p).X.
	p := &v
	p.X = 10
	fmt.Println(v.X) // 10: p punta a v

	// Le struct sono valori: l'assegnazione copia tutti i campi.
	w := v
	w.Y = 99
	fmt.Println(v, w) // v non cambia

	var zero Vertex
	fmt.Println(zero) // {0 0}
}
