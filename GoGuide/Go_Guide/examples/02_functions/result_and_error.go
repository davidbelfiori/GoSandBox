// result_and_error.go - Convenzione risultato + errore (slide 13, 15).
//
// La slide usa os.ReadFile("example.txt"). Qui il file viene prima creato in
// una directory temporanea, così l'esempio è autonomo e non dipende da file
// preesistenti; poi si prova anche un file inesistente per mostrare il ramo
// di errore. Il messaggio dell'errore dipende dal sistema operativo, quindi
// qui viene stampato solo il fatto che l'errore sia avvenuto.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readAndPrint(name string) {
	data, err := os.ReadFile(name)
	// Per convenzione l'errore è l'ultimo valore restituito e nil
	// significa "nessun errore". Va controllato subito.
	if err != nil {
		fmt.Println("Error: lettura fallita, err != nil")
		return
	}
	fmt.Println("File contents:", string(data))
}

func main() {
	dir, err := os.MkdirTemp("", "goguide")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer os.RemoveAll(dir) // pulizia al termine di main

	path := filepath.Join(dir, "example.txt")
	if err := os.WriteFile(path, []byte("ciao dal file"), 0o644); err != nil {
		fmt.Println("Error:", err)
		return
	}

	readAndPrint(path)                         // caso di successo
	readAndPrint(filepath.Join(dir, "no.txt")) // caso di errore
}
