package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Printf("Rilevato sistema operativo: %s\n", os)

	// Mappa per associare i sistemi operativi alle funzioni (print da eliminare)
	osFunctions := map[string]func(){
		"windows": runWindowsPath,
		"linux":   runLinuxPath,
	}

	// Esegui la funzione corrispondente al sistema operativo (print da eliminare)
	if fn, exists := osFunctions[os]; exists {
		fn()
	} else {
		fmt.Println("Sistema operativo non supportato!")
	}
}

func runWindowsPath() {
	fmt.Println("Eseguendo il percorso per Windows...")
	// Logica specifica per Windows (print da eliminare)
}

func runLinuxPath() {
	fmt.Println("Eseguendo il percorso per Linux...")
	// Logica specifica per Linux (print da eliminare)
}
