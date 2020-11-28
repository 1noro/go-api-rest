package model

// ProductInfo son los detalles extra de cada producto
type ProductInfo struct {
    Description string `json:"description"`
    Price int `json:"price"`
    AvailableAmount int `json:"available_amount"`
}