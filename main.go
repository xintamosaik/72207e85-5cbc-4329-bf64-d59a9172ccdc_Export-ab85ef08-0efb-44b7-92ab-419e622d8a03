package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	fileContents := getFileContents("index.html")
	fmt.Println(fileContents)

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
	switch r.Method {
	case http.MethodGet:
		// Handle GET request
		response := Response{Message: "This is a GET request"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	case http.MethodPost:
		// Handle POST request
		response := Response{Message: "This is a POST request"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
