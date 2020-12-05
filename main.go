package main

import (
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
    router.HandleFunc("/{username}/reserve/{reference}", resources.PostReserve).Queries("passwordSha", "{passwordSha}").Methods("POST")
    router.HandleFunc("/{username}/reserve/{reference}", resources.DeleteReserve).Queries("passwordSha", "{passwordSha}").Methods("DELETE")
    router.HandleFunc("/{username}/login", resources.CheckLogin).Queries("passwordSha", "{passwordSha}").Methods("GET")
    router.HandleFunc("/info/{text}", resources.GetURLInfo).Methods("GET")
    http.ListenAndServe(":8080", router)
}

