// methods.go - Metodi su tipi concreti (slide 41).
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Un metodo è una funzione con un receiver, scritto tra "func" e il nome.
// Con receiver per valore (v Vertex) il metodo lavora su una COPIA.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Con receiver puntatore (v *Vertex) il metodo può modificare l'originale.
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// Errore comune: receiver per valore, la modifica agisce sulla copia
// e va persa.
func (v Vertex) ScaleWrong(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 5

	v.ScaleWrong(10)
	fmt.Println("dopo ScaleWrong:", v) // invariato

	// v è una variabile indirizzabile: Go traduce v.Scale(10) in
	// (&v).Scale(10).
	v.Scale(10)
	fmt.Println("dopo Scale:", v, v.Abs())
}
