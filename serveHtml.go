package main

// how to use:
// go run serveHtml.go
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the file server to serve static files
	// fs := http.FileServer(http.Dir("./index.html"))
	// same as above but handle the error
	fs := http.FileServer(http.Dir("./"))
	if fs == nil {
		fmt.Println("Error reading file")
		return
	}

	// Register the file server handler and specify the URL path
	http.Handle("/", fs)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}	
