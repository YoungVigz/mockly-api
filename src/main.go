package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func generateMockData(w http.ResponseWriter, r *http.Request) {
	// Przykładowe użycie Mockly-CLI
	cmd := exec.Command("mockly", "generate", "-s", "./schemas/schema.json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		http.Error(w, "Błąd uruchamiania Mockly xddd", http.StatusInternalServerError)
		log.Println("Błąd:", err)
		return
	}

	fmt.Fprintln(w, "Mock data wygenerowane! okeya")
}

func main() {
	http.HandleFunc("/generate", generateMockData)

	fmt.Println("Serwer API uruchomiony na porcie 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
