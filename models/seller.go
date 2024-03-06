package models

type Seller struct {
	NmId             int    `json:"nmId"`
	SupplierId       int    `json:"supplierId"`       // id Поставщика
	SupplierName     string `json:"supplierName"`     // Краткое наименование поставщика
	SupplierFullName string `json:"supplierFullName"` // Полное наименование поставщика
	Inn              string `json:"inn"`              // ИНН
	Ogrnip           string `json:"ogrnip"`           // ОГРНИП (Основной Государственный Регистрационный Номер Индивидуального Предпринимателя)
	CountryLocCode   string `json:"countryLocCode"`   // Код страны (ru-RU)
	TaxpayerCode     string `json:"taxpayerCode"`     // Код налогоплательщика
	Trademark        string `json:"trademark"`        // Товарный знак
	Rv               int64  `json:"rv"`               // хз
}
