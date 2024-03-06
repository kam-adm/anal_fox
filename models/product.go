package models

import "time"

type Product struct {
	Id                    int       `json:"id"`
	Rating                int       `json:"rating"`                // Рэйтинг?
	Valuation             string    `json:"valuation"`             // Оценка
	ValuationToHundredths string    `json:"valuationToHundredths"` // Оценка до сотых
	FeedbacksCount        int       `json:"feedbacksCount"`        // Колличество отзывов
	UpdateDate            time.Time `json:"updateDate"`            // Дата обновления
	RatingIsInvisible     bool      `json:"ratingIsInvisible"`     // Видимость рейтинга?
	RegistrationDate      time.Time `json:"registrationDate"`      // Дата регистрации
	SaleItemQuantity      int       `json:"saleItemQuantity"`      // Товаров продано
	HasLogo               bool      `json:"hasLogo"`               // Имеет логотип
	HasBanner             bool      `json:"hasBanner"`             // Имеет баннер
	DefectPercent         int       `json:"defectPercent"`         // Процент дефектов
	DeliveryDuration      int       `json:"deliveryDuration"`      // Продолжительность доставки?
	SuppRatio             int       `json:"suppRatio"`             // хз
	RatioMarkSupp         int       `json:"ratioMarkSupp"`         // хз
	IsPremium             bool      `json:"isPremium"`             // Премиум?
}
