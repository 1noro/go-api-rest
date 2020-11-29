package main

import (
    // "fmt"
    "net/http"
    "./resources"
    "github.com/gorilla/mux" // go get -u github.com/gorilla/mux
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/health-check", resources.GetHealth).Methods("GET")
    router.HandleFunc("/product", resources.GetProducts).Methods("GET")
    router.HandleFunc("/product/{reference}", resources.GetFullProduct).Methods("GET")
    router.HandleFunc("/{username}/reserve", resources.GetReserves).Queries("passwordSha", "{passwordSha}").Methods("GET")
    http.ListenAndServe(":8080", router)
}

