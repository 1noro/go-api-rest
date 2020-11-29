package dataprovider

import (
    "../model"
)

// DataProvider interfaz estandar para develover los datos de este API REST
type DataProvider interface {
    GetProducts() []model.Product
    GetFullProduct(reference string) model.Product
    GetReserves(username string, passwordSha string) []model.Reserve
    PostReserve(reference string, username string, passwordSha string)
    DeleteReserve(reference string, username string, passwordSha string)
    CheckLogin(username string, passwordSha string) model.JSONHTTPResponse
}