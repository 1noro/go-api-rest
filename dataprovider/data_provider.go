package dataprovider

import (
    "../model"
)

// DataProvider interfaz estandar para develover los datos de este API REST
type DataProvider interface {
    GetProducts() ([]model.Product, int)
    GetFullProduct(reference string) (model.Product, int)
    GetReserves(username string, passwordSha string) ([]model.Reserve, int)
    PostReserve(reference string, username string, passwordSha string) int
    DeleteReserve(reference string, username string, passwordSha string) int
    CheckLogin(username string, passwordSha string) (model.JSONHTTPResponse, int)
}