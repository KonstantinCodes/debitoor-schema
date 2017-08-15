package debitoor

type PaymentaccountupdateV1 struct {
	AccountName string `json:"accountName"`
	AmountDelta float64 `json:"amountDelta"`
	Archived bool `json:"archived,omitempty"`
	Currency string `json:"currency,omitempty"`
	Id string `json:"id,omitempty"`
	IntegrationType string `json:"integrationType"`
	PaymentType string `json:"paymentType,omitempty"`
}