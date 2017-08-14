package debitoor

type PaymentAccountV1Sync struct {
	AccountName string `json:"accountName"`
	Archived bool `json:"archived,omitempty"`
	Balance float64 `json:"balance"`
	BankName string `json:"bankName"`
	Currency string `json:"currency"`
	Id string `json:"id"`
	IntegrationType string `json:"integrationType"`
	LastUpdateDate string `json:"lastUpdateDate,omitempty"`
	PaymentType string `json:"paymentType"`
	TransactionsCount float64 `json:"transactionsCount,omitempty"`
}