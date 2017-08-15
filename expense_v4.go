package debitoor

type ExpenseV4AssetDepreciation struct {
	AccumulatedDepreciation float64 `json:"accumulatedDepreciation,omitempty"`
	BookValue float64 `json:"bookValue,omitempty"`
	DepreciationCost float64 `json:"depreciationCost"`
	DepreciationDate string `json:"depreciationDate,omitempty"`
}

type ExpenseV4Expense struct {
	AttachmentOverQuota bool `json:"attachmentOverQuota,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	CreatedFromEmail bool `json:"createdFromEmail,omitempty"`
	Currency string `json:"currency,omitempty"`
	Date string `json:"date"`
	DeletedDate string `json:"deletedDate,omitempty"`
	File *ExpenseV4File `json:"file,omitempty"`
	FileId string `json:"fileId,omitempty"`
	Id string `json:"id,omitempty"`
	IncomeTaxDeductionAmount float64 `json:"incomeTaxDeductionAmount,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Lines []ExpenseV4Lines `json:"lines"`
	Links []ExpenseV4Links `json:"links,omitempty"`
	Number int `json:"number,omitempty"`
	Payments []ExpenseV4Payments `json:"payments,omitempty"`
	SupplierCiNumber string `json:"supplierCiNumber,omitempty"`
	SupplierCountry string `json:"supplierCountry,omitempty"`
	SupplierId string `json:"supplierId,omitempty"`
	SupplierName string `json:"supplierName,omitempty"`
	TaxGroups []ExpenseV4TaxGroups `json:"taxGroups,omitempty"`
	TotalGrossAmount float64 `json:"totalGrossAmount,omitempty"`
	TotalNetAmount float64 `json:"totalNetAmount,omitempty"`
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	UnpaidAmount float64 `json:"unpaidAmount,omitempty"`
}

type ExpenseV4File struct {
	FileName string `json:"fileName,omitempty"`
	Id string `json:"id,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	SizeBytes float64 `json:"sizeBytes,omitempty"`
	ThumbnailsUrl string `json:"thumbnailsUrl,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
}

type ExpenseV4Lines struct {
	AssetDepreciation []ExpenseV4AssetDepreciation `json:"assetDepreciation,omitempty"`
	AssetLifetime int `json:"assetLifetime,omitempty"`
	AssetSalvageValue float64 `json:"assetSalvageValue,omitempty"`
	CategoryId string `json:"categoryId,omitempty"`
	CategoryType string `json:"categoryType,omitempty"`
	Description string `json:"description,omitempty"`
	GrossAmount float64 `json:"grossAmount,omitempty"`
	IncomeTaxDeductionAmount float64 `json:"incomeTaxDeductionAmount,omitempty"`
	NetAmount float64 `json:"netAmount,omitempty"`
	PriceDisplayType string `json:"priceDisplayType,omitempty"`
	ReverseCharge *ExpenseV4ReverseCharge `json:"reverseCharge,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxRate float64 `json:"taxRate"`
}

type ExpenseV4Links struct {
	Amount float64 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Date string `json:"date,omitempty"`
	LinkId string `json:"linkId"`
	Number string `json:"number,omitempty"`
	Type string `json:"type"`
	UnpaidAmountChange float64 `json:"unpaidAmountChange,omitempty"`
}

type ExpenseV4Payments struct {
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

type ExpenseV4ReverseCharge struct {
	ConsumptionOrResale string `json:"consumptionOrResale"`
	ProductOrService string `json:"productOrService"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxRate float64 `json:"taxRate"`
}

type ExpenseV4TaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}