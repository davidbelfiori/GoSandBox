// defer.go - Esecuzione differita (slide 26).
package main

import "fmt"

func main() {
	// La chiamata viene registrata ora (argomenti valutati ora) ma eseguita
	// solo quando main sta per terminare.
	defer fmt.Println("world")
	fmt.Println("hello")
}
