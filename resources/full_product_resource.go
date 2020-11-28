package resources

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "../dataprovider"
)

// GetFullProduct devuelve el detalle de un producto
func GetFullProduct(responseWriter http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    reference := params["reference"]
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    responseWriter.Header().Set("Content-Type", "application/json")
    json.NewEncoder(responseWriter).Encode(dataProvider.GetFullProduct(reference))
}