package dataprovider

import (
    "../model"
)

// DataProvider interfaz estandar para develover los datos de este API REST
type DataProvider interface {
    GetProducts() []model.Product
    GetFullProduct(reference string) model.Product
}