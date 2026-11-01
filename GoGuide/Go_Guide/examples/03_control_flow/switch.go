// switch.go - switch senza fallthrough implicito (slide 25).
package main

import "fmt"

func describe(day string) {
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Almost weekend!")
	case "Saturday", "Sunday": // più valori nello stesso case
		fmt.Println("Weekend!")
	default:
		fmt.Println("Giorno qualunque")
	}
}

func main() {
	for _, d := range []string{"Monday", "Friday", "Sunday", "Wednesday"} {
		describe(d)
	}

	// I case non devono essere costanti: possono essere espressioni.
	// Senza espressione dopo "switch" equivale a "switch true": viene
	// eseguito il primo case la cui condizione è vera.
	n := 15
	switch {
	case n < 10:
		fmt.Println("piccolo")
	case n < 100:
		fmt.Println("medio")
	default:
		fmt.Println("grande")
	}
}
