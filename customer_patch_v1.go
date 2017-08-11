package debitoor

type Customer struct {
	Address string `json:"address,omitempty"`
	CiNumber string `json:"ciNumber,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	CountryName string `json:"countryName,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	CustomPaymentTermsDays float64 `json:"customPaymentTermsDays,omitempty"`
	Email string `json:"email,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Id string `json:"id,omitempty"`
	IsArchived bool `json:"isArchived,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Name string `json:"name,omitempty"`
	Notes string `json:"notes,omitempty"`
	Number int `json:"number,omitempty"`
	PaymentTermsId int `json:"paymentTermsId,omitempty"`
	Phone string `json:"phone,omitempty"`
	VatNumber string `json:"vatNumber,omitempty"`
}