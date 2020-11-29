package resources

import (
    // "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// CheckLogin verifica la utenticidad del usuario y la contraseña
func CheckLogin(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    // fmt.Println(params)
    username := params["username"]
    passwordSha := params["passwordSha"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    json.NewEncoder(responseWriter).Encode(dataProvider.CheckLogin(username, passwordSha))
}