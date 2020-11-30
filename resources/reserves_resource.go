package resources

import (
    // "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// GetReserves devuelve la lista de reservas de un usuario en JSON
func GetReserves(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    // fmt.Println(params)
    username := params["username"]
    passwordSha := params["passwordSha"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    reserves, _ := dataProvider.GetReserves(username, passwordSha)
    json.NewEncoder(responseWriter).Encode(reserves)
}