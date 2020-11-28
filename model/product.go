package model

// Product se la definición resumida de un poroducto
type Product struct {
    Reference string `json:"reference"`
    Name string `json:"name"`
    ImagePath string `json:"image_path"`
    ShortDescription string `json:"short_description"`
    ProductInfo ProductInfo `json:"product_info"`
}