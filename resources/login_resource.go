package resources

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// CheckLogin verifica la utenticidad del usuario y la contraseña
func CheckLogin(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    username := params["username"]
    passwordSha := params["passwordSha"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    httpState := dataProvider.CheckLogin(username, passwordSha)
    if httpState == 200 {
        json.NewEncoder(responseWriter).Encode(createOkResponse())
    } else if httpState == 401 {
        responseWriter.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(responseWriter).Encode(createUnauthorizedResponse())
    } else if httpState == 404 {
        responseWriter.WriteHeader(http.StatusNotFound)
        json.NewEncoder(responseWriter).Encode(createNotFoundResponse())
    } else if httpState == 500 {
        responseWriter.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(responseWriter).Encode(createInternalServerErrorResponse())
    }
}