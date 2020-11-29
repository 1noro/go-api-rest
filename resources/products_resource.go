package resources

import (
    "net/http"
    "encoding/json"
    "../dataprovider"
)

// GetProducts devuelve una lista JSON con los resumenes de los productos para el catálogo
func GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.Header().Set("Content-Type", "application/json")
    var dataProvider dataprovider.DataProvider
    dataProvider = dataprovider.GetDataProvider()
    json.NewEncoder(responseWriter).Encode(dataProvider.GetProducts())
}