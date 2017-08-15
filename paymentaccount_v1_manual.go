package debitoor

type PaymentAccountV1Manual struct {
	AccountName string `json:"accountName"`
	AmountDelta float64 `json:"amountDelta"`
	Archived bool `json:"archived,omitempty"`
	Balance float64 `json:"balance,omitempty"`
	Currency string `json:"currency"`
	Id string `json:"id,omitempty"`
	IntegrationType string `json:"integrationType"`
	PaymentType string `json:"paymentType"`
	TransactionsCount float64 `json:"transactionsCount,omitempty"`
}