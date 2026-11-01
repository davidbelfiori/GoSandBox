// helloworld.go - Primo programma Go (slide 8-9).
//
// Ogni file inizia con la dichiarazione del package. Il package "main" è
// speciale: dice al toolchain di produrre un eseguibile, il cui punto di
// ingresso è la funzione main().
package main

// "fmt" fornisce l'I/O formattato (Println, Printf, ...).
import "fmt"

func main() {
	// Le stringhe Go sono UTF-8 nativamente: nessuna configurazione
	// particolare serve per stampare caratteri non ASCII.
	fmt.Println("Hello, 世界")
}
