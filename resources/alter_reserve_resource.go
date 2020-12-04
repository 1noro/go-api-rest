package resources

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// PostReserve crea una reserva para un usuario
func PostReserve(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    username := params["username"]
    passwordSha := params["passwordSha"]
    reference := params["reference"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    httpState := dataProvider.PostReserve(reference, username, passwordSha)
    if httpState == 201 {
        json.NewEncoder(responseWriter).Encode(createCreatedResponse())
    } else if httpState == 401 {
        responseWriter.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(responseWriter).Encode(createUnauthorizedResponse())
    } else if httpState == 404 {
        responseWriter.WriteHeader(http.StatusNotFound)
        json.NewEncoder(responseWriter).Encode(createNotFoundResponse())
    } else if httpState == 409 {
        responseWriter.WriteHeader(http.StatusNotFound)
        json.NewEncoder(responseWriter).Encode(createConflictResponse())
    } else if httpState == 500 {
        responseWriter.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(responseWriter).Encode(createInternalServerErrorResponse())
    }
}

// DeleteReserve birra una reserva para un usuario
func DeleteReserve(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    username := params["username"]
    passwordSha := params["passwordSha"]
    reference := params["reference"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    httpState := dataProvider.DeleteReserve(reference, username, passwordSha)
    if httpState == 200 {
        json.NewEncoder(responseWriter).Encode(createOkResponseExtra("Deleted"))
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