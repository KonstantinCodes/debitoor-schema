package debitoor

type SupplierV1 struct {
	Address string `json:"address,omitempty"`
	CiNumber string `json:"ciNumber,omitempty"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	Email string `json:"email,omitempty"`
	Id string `json:"id,omitempty"`
	IsArchived bool `json:"isArchived,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Name string `json:"name"`
	Number int `json:"number,omitempty"`
	Phone string `json:"phone,omitempty"`
}