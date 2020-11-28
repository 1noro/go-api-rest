package main

import (
    // "fmt"
    "net/http"
    "./resources"
)

func main() {
    http.HandleFunc("/health-check", resources.GetHealth)
    http.HandleFunc("/product", resources.GetProducts)
    http.ListenAndServe(":8080", nil)
}

