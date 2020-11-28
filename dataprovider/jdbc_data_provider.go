package dataprovider

import (
    "../model"
)

// JDBCDataProvider conexión entre el API REST y la base de datos
type JDBCDataProvider struct {}

// GetProducts devuelve una lista de productos para el catálogo
func (jdbcDataProvider JDBCDataProvider) GetProducts() []model.Product {
    p1 := model.Product{Reference:"R-001", Name:"Truck1", ImagePath:"img.jpg", ShortDescription:"aaa"}
    p2 := model.Product{Reference:"R-002", Name:"Truck2", ImagePath:"img.jpg", ShortDescription:"aaa"}
    out := []model.Product{p1, p2}
    return out
}