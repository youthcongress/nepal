package server

import (
	"fmt"
	"net/http"
)

// Import routes
func StartServer() {
	// Define the directory to serve static files
	publicDir := "./public"
	fs := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fs)

	// Register routes
	SetupRoutes()

	// Define the server port
	port := 8080

	// Start the server
	fmt.Printf("🚀 Server running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("❌ Server failed to start:", err)
	}
}
