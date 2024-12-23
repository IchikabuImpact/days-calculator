package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Response represents the JSON response structure
type Response struct {
	Date string `json:"date"`
}

// calculateDate computes the date `daysAgo` days before today
func calculateDate(daysAgo int) string {
	currentDate := time.Now()
	pastDate := currentDate.AddDate(0, 0, -daysAgo)
	return pastDate.Format("2006/01/02")
}

// handleDaysCalculator handles HTTP requests for the days calculator
func handleDaysCalculator(w http.ResponseWriter, r *http.Request) {
	daysQuery := r.URL.Query().Get("days")
	if daysQuery == "" {
		http.Error(w, "Missing 'days' parameter", http.StatusBadRequest)
		return
	}

	daysAgo, err := strconv.Atoi(daysQuery)
	if err != nil {
		http.Error(w, "'days' parameter must be a valid integer", http.StatusBadRequest)
		return
	}

	result := calculateDate(daysAgo)
	json.NewEncoder(w).Encode(Response{Date: result})
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found. Using default settings.")
	}

	// Get the PORT from environment variables or use the default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// If CLI arguments are provided, run in CLI mode
	if len(os.Args) > 1 {
		daysAgo, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: Argument must be a valid integer")
			os.Exit(1)
		}
		fmt.Println(calculateDate(daysAgo))
		return
	}

	// Start the HTTP server
	http.HandleFunc("/api/calculate", handleDaysCalculator)
	fmt.Printf("Server started at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

