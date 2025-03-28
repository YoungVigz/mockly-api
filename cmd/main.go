package main

import (
	"fmt"
	"log"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Database initialization and connection test
	err := database.DatabaseInit()

	if err != nil {
		log.Fatal("Unable to connect to database!")
	}

	fmt.Println("Connected to database")

	// Running API
	api := gin.Default()
	routes.RegisterRoutes(api)
	api.Run(":8080")
}

/*
func generateMockData(w http.ResponseWriter, r *http.Request) {
	// Przykładowe użycie Mockly-CLI
	cmd := exec.Command("mockly", "generate", "-s", "./schemas/schema.json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		http.Error(w, "Błąd uruchamiania Mockly", http.StatusInternalServerError)
		log.Println("Błąd:", err)
		return
	}

	fmt.Fprintln(w, "Mock data wygenerowane!")
}
*/
