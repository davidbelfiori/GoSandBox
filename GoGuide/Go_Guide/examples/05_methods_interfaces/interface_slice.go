// interface_slice.go - Polimorfismo: una collezione eterogenea di forme.
//
// Esempio aggiuntivo (non presente nelle slide) che riusa solo concetti già
// visti: interfacce implicite, metodi e range su slice.
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct{ side float64 }
type circle struct{ radius float64 }

func (s square) area() float64 { return s.side * s.side }
func (c circle) area() float64 { return math.Pi * c.radius * c.radius }

// totalArea somma le aree di tipi diversi: a runtime ogni chiamata
// s.area() viene risolta sul tipo concreto contenuto nell'interfaccia.
func totalArea(shapes []shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.area()
	}
	return total
}

func main() {
	shapes := []shape{square{2}, circle{1}, square{3}}
	fmt.Printf("area totale = %.4f\n", totalArea(shapes))
}
