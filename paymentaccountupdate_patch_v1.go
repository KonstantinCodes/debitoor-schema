package debitoor

type PaymentaccountupdatePatchV1 struct {
	AccountName string `json:"accountName,omitempty"`
	AmountDelta float64 `json:"amountDelta,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Currency string `json:"currency,omitempty"`
	Id string `json:"id,omitempty"`
	IntegrationType string `json:"integrationType,omitempty"`
	PaymentType string `json:"paymentType,omitempty"`
}