package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type jSONText struct {
    Text string `json:"text"`
}

func getHealth(responseWriter http.ResponseWriter, request *http.Request) {
    fmt.Print(request)
    jsonResponse := jSONText{"Health Check OK"}
    responseWriter.Header().Set("Content-Type", "application/json")
    json.NewEncoder(responseWriter).Encode(jsonResponse)
}

func main() {
    http.HandleFunc("/health-check", getHealth)
    http.ListenAndServe(":8080", nil)
}

