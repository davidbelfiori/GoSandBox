// types.go - Tipi base, conversioni esplicite e inferenza (slide 18).
package main

import "fmt"

func main() {
	// Tipi base.
	var (
		b  bool    = true
		s  string  = "go"
		i  int     = 42
		u  uint    = 7
		f2 float32 = 1.5
		f6 float64 = 2.25
	)
	fmt.Println(b, s, i, u, f2, f6)

	// Conversione esplicita: obbligatoria tra tipi diversi (a differenza
	// di C, dove int -> double avviene implicitamente).
	var n int = 42
	var f float64 = float64(n)
	fmt.Printf("%d -> %g\n", n, f)

	// Inferenza: j ha lo stesso tipo di k (int).
	var k int
	j := k
	fmt.Printf("tipo di j: %T\n", j)

	// Il tipo dedotto da una costante letterale dipende dalla sua forma.
	x := 3.14 // float64
	y := 'a'  // rune (alias di int32)
	fmt.Printf("%T %T\n", x, y)

	// Non si può sommare int e float64 senza conversione:
	//   fmt.Println(n + f)   // errore di compilazione: mismatched types
	fmt.Println(float64(n) + f)
}
