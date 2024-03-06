package models

type Request struct {
	Seller       Seller         `json:"seller"`
	Compositions []Composition  `json:"composition"`
	Product      Product        `json:"product"`
	PriceHistory []PriceHistory `json:"priceHistory"`
}
