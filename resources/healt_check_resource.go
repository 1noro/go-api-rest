package resources

import (
    "net/http"
    "encoding/json"
    "../model"
)

// GetHealth devuelve un JSONText que demuestra que el servidor está funcionando
func GetHealth(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    jsonResponse := model.JSONText{Text:"Health Check OK"}
    json.NewEncoder(responseWriter).Encode(jsonResponse)
}