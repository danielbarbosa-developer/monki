package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Monki API is running 🚀")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting Monki server on port 8080...")
    http.ListenAndServe(":9000", nil)
}
