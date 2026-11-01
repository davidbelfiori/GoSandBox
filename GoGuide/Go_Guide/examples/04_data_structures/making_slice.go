// making_slice.go - Creare slice con make (slide 30-31).
package main

import "fmt"

func main() {
	a := make([]int, 0, 5) // len(s)=0, cap(s)=5
	printSlice("a", a)

	b := make([]int, 5) // len=5, cap=5: cinque zeri
	printSlice("b", b)

	c := make([]int, 5, 10)
	printSlice("c", c)

	d := c[2:4] // len=2, cap=8 (si parte dall'indice 2 dell'array)
	printSlice("d", d)

	var nilSlice []int // slice nil: len=0, cap=0, ma utilizzabile con append
	printSlice("nil", nilSlice)
	fmt.Println("nilSlice == nil?", nilSlice == nil)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
