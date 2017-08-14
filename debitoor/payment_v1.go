package debitoor

type PaymentV1 struct {
	Amount float64 `json:"amount"`
	CategoryIds []string `json:"categoryIds,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	Currency string `json:"currency"`
	CustomerName string `json:"customerName,omitempty"`
	Description string `json:"description,omitempty"`
	ExpenseId string `json:"expenseId,omitempty"`
	Id string `json:"id,omitempty"`
	IncomeId string `json:"incomeId,omitempty"`
	IntegrationType string `json:"integrationType,omitempty"`
	InvoiceId string `json:"invoiceId,omitempty"`
	InvoiceNumber string `json:"invoiceNumber,omitempty"`
	MatchedPaymentAccountId string `json:"matchedPaymentAccountId,omitempty"`
	MatchedPaymentTransactionId string `json:"matchedPaymentTransactionId,omitempty"`
	PaymentAccountId string `json:"paymentAccountId"`
	PaymentDate string `json:"paymentDate,omitempty"`
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	PaymentType string `json:"paymentType,omitempty"`
	SupplierName string `json:"supplierName,omitempty"`
	Text string `json:"text,omitempty"`
}