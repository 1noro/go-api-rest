package dataprovider

import (
    "../model"
)

// MockDataProvider proveedor de datos de prueba
type MockDataProvider struct {}

// GetProducts devuelve una lista de productos para el catálogo
func (mockDataProvider MockDataProvider) GetProducts() []model.Product {
    product1 := model.Product{
        Reference : "R-001",
        Name : "BatTruck",
        ImagePath : "https://i.ibb.co/7K2H8k1/01.png",
        ShortDescription : "El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
    }
    pproduct2 := model.Product{
        Reference : "R-002",
        Name : "Bandit Truck V8",
        ImagePath : "https://i.ibb.co/F8mm32n/02.png",
        ShortDescription : "Uno de los mejores vehículos del catálogo. Gran estabilidad y capacidad de aplastar todo a su paso. Este vehículo fue...",
    }
    out := []model.Product{product1, pproduct2}
    return out
}

// GetFullProduct devuelve el detalle de un producto
func (mockDataProvider MockDataProvider) GetFullProduct(reference string) model.Product {
    info := model.ProductInfo{
        Description : "El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
        Price : 799999,
        AvailableAmount : 8,
    }
    product := model.Product{
        Reference : reference,
        Name : "BatTruck",
        ImagePath : "https://i.ibb.co/7K2H8k1/01.png",
        ShortDescription : "El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
        ProductInfo : info,
    }
    return product
}