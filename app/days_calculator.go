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

// now returns the current time. It is defined as a variable so tests can
// override it for deterministic output.
var now = time.Now

// dateNDaysAgo returns the date `daysAgo` days prior to today.
func dateNDaysAgo(daysAgo int) string {
	currentDate := now()
	pastDate := currentDate.AddDate(0, 0, -daysAgo)
	return pastDate.Format("2006/01/02")
}

// daysCalculatorHandler handles HTTP requests for the days calculator.
func daysCalculatorHandler(w http.ResponseWriter, r *http.Request) {
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

	result := dateNDaysAgo(daysAgo)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Response{Date: result}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
		fmt.Println(dateNDaysAgo(daysAgo))
		return
	}

	// Start the HTTP server
	http.HandleFunc("/api/calculate", daysCalculatorHandler)
	fmt.Printf("Server started at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
