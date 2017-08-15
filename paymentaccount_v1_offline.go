package debitoor

type PaymentAccountOfflineV1Amount struct {
	Column float64 `json:"column"`
	DecimalSeparator string `json:"decimalSeparator"`
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type PaymentAccountOfflineV1AutoMatch struct {
	Disabled bool `json:"disabled,omitempty"`
}

type PaymentAccountOfflineV1Date struct {
	Column float64 `json:"column"`
	Format string `json:"format,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type PaymentAccountOfflineV1Mapping struct {
	Amount *PaymentAccountOfflineV1Amount `json:"amount,omitempty"`
	Currency string `json:"currency"`
	Date *PaymentAccountOfflineV1Date `json:"date,omitempty"`
	OtherParty *PaymentAccountOfflineV1OtherParty `json:"otherParty,omitempty"`
	Text *PaymentAccountOfflineV1Text `json:"text,omitempty"`
}

type PaymentAccountOfflineV1OfflinePaymentAccountMetadataToBeSavedAsAPartOfUserSettings struct {
	AccountName string `json:"accountName"`
	Archived bool `json:"archived,omitempty"`
	AutoMatch *PaymentAccountOfflineV1AutoMatch `json:"autoMatch,omitempty"`
	Balance float64 `json:"balance"`
	BankId string `json:"bankId"`
	BankName string `json:"bankName,omitempty"`
	Currency string `json:"currency"`
	Id string `json:"id,omitempty"`
	IntegrationType string `json:"integrationType"`
	LastUpdateDate string `json:"lastUpdateDate,omitempty"`
	Mapping *PaymentAccountOfflineV1Mapping `json:"mapping,omitempty"`
	PaymentType string `json:"paymentType"`
	StatementId string `json:"statementId"`
	TransactionsCount float64 `json:"transactionsCount,omitempty"`
}

type PaymentAccountOfflineV1OtherParty struct {
	Column float64 `json:"column"`
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type PaymentAccountOfflineV1Text struct {
	Column float64 `json:"column"`
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}