package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Hello, %s!\n", "World")
	fmt.Print("Hello, World! \n")

	z, k := swap("bella", "raga")
	fmt.Println(z, k)
	fmt.Println(add(42, 13))

}

/*
fmt.Print stampa i valori uno dopo l’altro, senza andare a capo. Aggiunge uno spazio tra due operandi solo se nessuno dei due è una stringa.
//fmt.Println stampa i valori separati sempre da spazi e aggiunge un a capo alla fine.
//fmt.Printf usa una stringa di formato con dei segnaposto (verbi come %d, %s, %f, %v). Non aggiunge né spazi né a capo, quindi il \n va scritto esplicitamente.
// nomeFunzione(parametro1 tipo, parametro2 tipo) tipoRitorno{  // sembra una follia ma ce lo teniamo
*/
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
