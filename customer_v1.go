package debitoor

type CustomerV1 struct {
	Address string `json:"address,omitempty"`
	CiNumber string `json:"ciNumber,omitempty"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	CustomPaymentTermsDays float64 `json:"customPaymentTermsDays,omitempty"`
	Email string `json:"email,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Id string `json:"id,omitempty"`
	IsArchived bool `json:"isArchived,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Name string `json:"name"`
	Notes string `json:"notes,omitempty"`
	Number int `json:"number,omitempty"`
	PaymentTermsId int `json:"paymentTermsId"`
	Phone string `json:"phone,omitempty"`
	VatNumber string `json:"vatNumber,omitempty"`
}