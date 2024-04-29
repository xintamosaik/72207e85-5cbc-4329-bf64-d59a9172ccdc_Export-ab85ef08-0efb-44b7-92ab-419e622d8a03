package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

func CSV2Map(csv string, header string) map[string]string {
	// Split the header string into a slice of strings
	headerSlice := strings.Split(header, ", ")

	// Split the CSV string into a slice of strings
	csvSlice := strings.Split(csv, ", ")

	// Create a map to store the CSV data
	csvMap := make(map[string]string)

	// Iterate over the header slice
	for i, header := range headerSlice {
		// Add the header and corresponding CSV value to the map
		csvMap[header] = csvSlice[i]
	}

	return csvMap
}
func main() {

	fs := http.FileServer(http.Dir("./"))
	if fs == nil {
		fmt.Println("Error reading file")
		return
	}
	// (index):13 GET http://localhost:3000/api net::ERR_CONNECTION_REFUSED
	// Why? Because we haven't implemented the /api route yet. Let's do that now.

	// Register the file server handler and specify the URL path

	// Start the HTTP server
	http.Handle("/", fs)
	http.HandleFunc("/api", handleRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func getFileContents(filePath string) string {

	// Read the file
	content, err := os.ReadFile(filePath)

	// Handle the error if we found one
	if err != nil {
		fmt.Println("Error reading file")
		return ""
	}
	return string(content)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// column names
	// exampleCSVHeader := "description, start_date, end_date, amount, recurring, mon, tue, wed, thu, fri, sat, sun, weekly, biweekly, monthly"
	// comma separated values
	exampleCSV := "buy new glasses, 2021-06-01, 2021-06-30, 1000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1"
	// CSVMap := CSV2Map(exampleCSV, exampleCSVHeader)

	// Handle GET request
	response := Response{Message: exampleCSV}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
