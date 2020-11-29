package model

// Reserve devuelve los datos de un producto reservado en una fecha
type Reserve struct {
    Product Product `json:"product"`
    ReserveDate int64 `json:"reserve_date"`
}