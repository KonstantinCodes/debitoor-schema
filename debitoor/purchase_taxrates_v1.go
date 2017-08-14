package debitoor

type PurchaseTaxratesV1BaseTax struct {
	DefaultRate float64 `json:"defaultRate,omitempty"`
	Map *PurchaseTaxratesV1Map `json:"map,omitempty"`
	Rates []float64 `json:"rates,omitempty"`
}

type PurchaseTaxratesV1Map struct {
}

type PurchaseTaxratesV1ReverseCharge struct {
	Rates []float64 `json:"rates,omitempty"`
}

type PurchaseTaxratesV1Root struct {
	BaseTax *PurchaseTaxratesV1BaseTax `json:"baseTax,omitempty"`
	CategoryId string `json:"categoryId,omitempty"`
	Date string `json:"date,omitempty"`
	MapFromCategoryId string `json:"mapFromCategoryId,omitempty"`
	MapFromDate string `json:"mapFromDate,omitempty"`
	MapFromSupplierCountry string `json:"mapFromSupplierCountry,omitempty"`
	ReverseCharge *PurchaseTaxratesV1ReverseCharge `json:"reverseCharge,omitempty"`
	SupplierCountry string `json:"supplierCountry,omitempty"`
}