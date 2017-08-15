package debitoor

type SalesTaxratesV1BaseTax struct {
	DefaultRate float64 `json:"defaultRate,omitempty"`
	Map *SalesTaxratesV1Map `json:"map,omitempty"`
	ProductServiceSplit bool `json:"productServiceSplit,omitempty"`
	Rates []float64 `json:"rates,omitempty"`
}

type SalesTaxratesV1Map struct {
}

type SalesTaxratesV1Root struct {
	BaseTax *SalesTaxratesV1BaseTax `json:"baseTax,omitempty"`
	CustomerCountry string `json:"customerCountry,omitempty"`
	Date string `json:"date,omitempty"`
	DomesticToIntraEu bool `json:"domesticToIntraEu,omitempty"`
	MapFromCustomerCountry string `json:"mapFromCustomerCountry,omitempty"`
	MapFromDate string `json:"mapFromDate,omitempty"`
}