package main

import (
	"fmt"
	"net/http"
	"processus/api"
)

func main() {
	mux := http.NewServeMux()

	// Iterate over the map using range
	paths := api.GetProcessesConfigurationsPaths()
	for key, value := range paths {
		mux.HandleFunc(key, value)
	}
	fmt.Println("Serveur en cours d'exécution sur le port 8080...")
	handler := corsMiddleware(mux)

	http.ListenAndServe(":8080", handler)
}

// corsMiddleware is a middleware function to handle CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
