// interface.go - Interfacce implicite e polimorfismo (slide 42-45).
package main

import (
	"fmt"
	"math"
)

// Define the interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// Implement the interface on rectangular and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Per implementare un'interfaccia in Go basta implementare tutti i suoi
// metodi: nessuna dichiarazione "implements".

// Implement area and perim methods on rect
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implement area and perim methods on circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Se una variabile ha tipo interfaccia si possono chiamare i metodi
// dell'interfaccia: measure funziona con qualsiasi tipo che implementi
// geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// rect e circle soddisfano entrambe geometry, quindi possono essere
	// passate a measure.
	measure(r)
	measure(c)
}
