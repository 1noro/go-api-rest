package resources

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../model"
)

// GetURLInfo devuelve un JSONText con la URL que ha solicitado este recurso
func GetURLInfo(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    text := params["text"]
    jsonResponse := model.JSONText{Text:text}
    json.NewEncoder(responseWriter).Encode(jsonResponse)
}