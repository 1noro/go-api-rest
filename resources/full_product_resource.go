package resources

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// GetFullProduct devuelve el detalle de un producto
func GetFullProduct(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    reference := params["reference"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    product, httpState := dataProvider.GetFullProduct(reference)
    if httpState == 200 {
        json.NewEncoder(responseWriter).Encode(product)
    } else if httpState == 404 {
        responseWriter.WriteHeader(http.StatusNotFound)
        json.NewEncoder(responseWriter).Encode(createNotFoundResponse())
    } else if httpState == 500 {
        responseWriter.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(responseWriter).Encode(createInternalServerErrorResponse())
    }
}