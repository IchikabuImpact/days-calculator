package main

import (
	"encoding/json"
	"flag"
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

// runCLI parses command line arguments and returns the calculated date.
// It expects either a -days flag or a positional argument providing the number
// of days. When no arguments are supplied it returns an error.
func runCLI(args []string) (string, error) {
	fs := flag.NewFlagSet("days-calculator", flag.ContinueOnError)
	days := fs.Int("days", -1, "number of days ago")
	if err := fs.Parse(args); err != nil {
		return "", err
	}

	var d int
	if *days >= 0 {
		d = *days
	} else if fs.NArg() > 0 {
		var err error
		d, err = strconv.Atoi(fs.Arg(0))
		if err != nil {
			return "", fmt.Errorf("argument must be a valid integer")
		}
	} else {
		return "", fmt.Errorf("no days provided")
	}

	return dateNDaysAgo(d), nil
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

	// If CLI arguments are provided, run in CLI mode using the flag package
	if len(os.Args) > 1 {
		result, err := runCLI(os.Args[1:])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
		return
	}

	// Start the HTTP server
	http.HandleFunc("/api/calculate", daysCalculatorHandler)
	fmt.Printf("Server started at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
