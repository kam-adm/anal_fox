package models

type PriceHistory struct {
	Dt    int `json:"dt"` // Дата изменения цены
	Price struct {
		RUB int `json:"RUB"` // Цена
	} `json:"price"`
}

// Не отображает информацию о текущей цене
