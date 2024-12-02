package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Rileva il sistema operativo
	os := runtime.GOOS
	fmt.Printf("Rilevato sistema operativo: %s\n", os)

	// Percorsi specifici per ogni sistema operativo
	switch os {
	case "windows":
		runWindowsPath()
	case "linux":
		runLinuxPath()
	default:
		fmt.Println("Sistema operativo non supportato!")
	}
}

// Funzione per il percorso Windows
func runWindowsPath() {
	fmt.Println("Eseguendo il percorso per Windows...")
	// Logica specifica per Windows
}

// Funzione per il percorso Linux
func runLinuxPath() {
	fmt.Println("Eseguendo il percorso per Linux...")
	// Logica specifica per Linux
}
