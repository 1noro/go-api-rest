package dataprovider

import (
    // "time"
    "../model"
)

// MockDataProvider proveedor de datos de prueba
type MockDataProvider struct {}

// GetProducts devuelve una lista de productos para el catálogo
func (mockDataProvider MockDataProvider) GetProducts() []model.Product {
    product1 := model.Product{
        Reference:"R-001",
        Name:"BatTruck",
        ImagePath:"https://i.ibb.co/7K2H8k1/01.png",
        ShortDescription:"El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
    }
    pproduct2 := model.Product{
        Reference:"R-002",
        Name:"Bandit Truck V8",
        ImagePath:"https://i.ibb.co/F8mm32n/02.png",
        ShortDescription:"Uno de los mejores vehículos del catálogo. Gran estabilidad y capacidad de aplastar todo a su paso. Este vehículo fue...",
    }
    return []model.Product{product1, pproduct2}
}

// GetFullProduct devuelve el detalle de un producto
func (mockDataProvider MockDataProvider) GetFullProduct(reference string) model.Product {
    info := model.ProductInfo{
        Description:"El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
        Price:799999,
        AvailableAmount:8,
    }
    product := model.Product{
        Reference:reference,
        Name:"BatTruck",
        ImagePath:"https://i.ibb.co/7K2H8k1/01.png",
        ShortDescription:"El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
        ProductInfo:info,
    }
    return product
}

// GetReserves devuelve la lista de reservas de un usuario
func (mockDataProvider MockDataProvider) GetReserves(username string, passwordSha string) []model.Reserve {
    info := model.ProductInfo{
        Description:"El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
        Price:799999,
        AvailableAmount:8,
    }
    product := model.Product{
        Reference:"R-001",
        Name:"BatTruck",
        ImagePath:"https://i.ibb.co/7K2H8k1/01.png",
        ShortDescription:"El camión monstruo más heroico. Ganó la copa mundial TruckStarts 2017 contra el TruckJoker.",
        ProductInfo:info,
    }
    reserve := model.Reserve{
        Product:product,
        // ReserveDate:time.Now().Unix(),
    }
    return []model.Reserve{reserve}
}

// PostReserve crea una reserva nueva para un usuario
func (mockDataProvider MockDataProvider) PostReserve(reference string, username string, passwordSha string) {}

// DeleteReserve borra una reserva nueva para un usuario
func (mockDataProvider MockDataProvider) DeleteReserve(reference string, username string, passwordSha string) {}

// CheckLogin comprueba si el usuario y la contraseña son correctos
func (mockDataProvider MockDataProvider) CheckLogin(username string, passwordSha string) model.JSONHTTPResponse {
    return model.JSONHTTPResponse{HTTPResponse:model.HTTPResponse{Code:200, Description: "OK", ExtraText: "Login check OK"}}
}