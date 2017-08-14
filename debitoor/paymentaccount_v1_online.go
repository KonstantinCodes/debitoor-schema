package debitoor

type PaymentAccountOnlineV1AutoMatch struct {
	Disabled bool `json:"disabled,omitempty"`
}

type PaymentAccountOnlineV1Credentials struct {
}

type PaymentAccountOnlineV1 struct {
	AccountName string `json:"accountName,omitempty"`
	Archived bool `json:"archived,omitempty"`
	AutoMatch *PaymentAccountOnlineV1AutoMatch `json:"autoMatch,omitempty"`
	Balance float64 `json:"balance,omitempty"`
	BankId string `json:"bankId"`
	BankName string `json:"bankName,omitempty"`
	Credentials *PaymentAccountOnlineV1Credentials `json:"credentials,omitempty"`
	Currency string `json:"currency,omitempty"`
	Id string `json:"id"`
	IntegrationType string `json:"integrationType"`
	LastUpdateDate string `json:"lastUpdateDate,omitempty"`
	PaymentType string `json:"paymentType"`
	SavePin bool `json:"savePin,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	TransactionsCount float64 `json:"transactionsCount,omitempty"`
}