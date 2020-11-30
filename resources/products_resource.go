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
    products, httpState := dataProvider.GetProducts()
    if httpState == 200 {
        json.NewEncoder(responseWriter).Encode(products)
    } else if httpState == 500 {
        responseWriter.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(responseWriter).Encode(createInternalServerErrorResponse())
    }
}