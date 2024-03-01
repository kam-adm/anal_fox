package models

type Seller struct {
	NmId             int    `json:"nmId"` // ебола
	SupplierId       int    `json:"supplierId"`
	SupplierName     string `json:"supplierName"`
	SupplierFullName string `json:"supplierFullName"`
	Inn              string `json:"inn"`
	Ogrnip           string `json:"ogrnip"`
	CountryLocCode   string `json:"countryLocCode"`
	TaxpayerCode     string `json:"taxpayerCode"`
	Trademark        string `json:"trademark"`
	Rv               int64  `json:"rv"`
}

type Request struct {
	Seller       Seller        `json:"seller"`
	Compositions []Composition `json:"composition"`
}
