package model

// Product se la definición resumida de un poroducto
type Product struct {
    Reference string `json:"reference"`
    Name string `json:"name"`
    ImagePath string `json:"imagePath"` // `json:"image_path"`
    ShortDescription string `json:"shortDescription"` // `json:"short_description"`
    ProductInfo ProductInfo `json:"productInfo"` // `json:"product_info"`
}