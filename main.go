package main

import (
    // "fmt"
    "net/http"
    "./resources"
)

func main() {
    http.HandleFunc("/health-check", resources.GetHealth)
    http.ListenAndServe(":8080", nil)
}

