package debitoor

type ProductPatchV1 struct {
	CalculatedPrice string `json:"calculatedPrice,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	Description string `json:"description,omitempty"`
	GrossUnitSalesPrice float64 `json:"grossUnitSalesPrice,omitempty"`
	Id string `json:"id,omitempty"`
	IsArchived bool `json:"isArchived,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Name string `json:"name,omitempty"`
	NetUnitCostPrice float64 `json:"netUnitCostPrice,omitempty"`
	NetUnitSalesPrice float64 `json:"netUnitSalesPrice,omitempty"`
	ProductOrService string `json:"productOrService,omitempty"`
	Rate float64 `json:"rate,omitempty"`
	Sku string `json:"sku,omitempty"`
	TaxEnabled bool `json:"taxEnabled,omitempty"`
	UnitId string `json:"unitId,omitempty"`
	UnitName string `json:"unitName,omitempty"`
}