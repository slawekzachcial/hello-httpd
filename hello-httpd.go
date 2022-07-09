package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HELLO_HTTPD_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Greeting on port %s...\n", port)

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	greeted := r.URL.Path[1:]
	if greeted == "" {
		greeted = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", greeted)
}
