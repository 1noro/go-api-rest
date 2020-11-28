package resources

import (
	// "fmt"
    "net/http"
    "encoding/json"
    "../model"
)

// GetHealth devuelve un JSONText que demuestra que el servidor está funcionando
func GetHealth(responseWriter http.ResponseWriter, request *http.Request) {
    // fmt.Print(request)
    jsonResponse := model.JSONText{Text:"Health Check OK"}
    responseWriter.Header().Set("Content-Type", "application/json")
    json.NewEncoder(responseWriter).Encode(jsonResponse)
}