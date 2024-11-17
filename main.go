package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Define the server configuration
	server := &http.Server{
		Addr: ":8080", // Port to listen on
	}

	fmt.Println("Starting server on :8080")
	// Start the server
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
