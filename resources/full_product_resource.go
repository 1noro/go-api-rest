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
    json.NewEncoder(responseWriter).Encode(dataProvider.GetFullProduct(reference))
}