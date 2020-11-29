package resources

import (
    // "fmt"
    "net/http"
    // "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// PostReserve crea una reserva para un usuario
func PostReserve(responseWriter http.ResponseWriter, request *http.Request) {
    // responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    // fmt.Println(params)
    username := params["username"]
    passwordSha := params["passwordSha"]
    reference := params["reference"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    dataProvider.PostReserve(reference, username, passwordSha)
    // json.NewEncoder(responseWriter).Encode()
}

// DeleteReserve birra una reserva para un usuario
func DeleteReserve(responseWriter http.ResponseWriter, request *http.Request) {
    // responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    // fmt.Println(params)
    username := params["username"]
    passwordSha := params["passwordSha"]
    reference := params["reference"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    dataProvider.DeleteReserve(reference, username, passwordSha)
    // json.NewEncoder(responseWriter).Encode()
}