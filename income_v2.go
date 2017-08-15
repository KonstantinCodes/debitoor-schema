package debitoor

type IncomeV2AssetDepreciation struct {
	AccumulatedDepreciation float64 `json:"accumulatedDepreciation,omitempty"`
	BookValue float64 `json:"bookValue,omitempty"`
	DepreciationCost float64 `json:"depreciationCost"`
	DepreciationDate string `json:"depreciationDate,omitempty"`
}

type IncomeV2File struct {
	FileName string `json:"fileName,omitempty"`
	Id string `json:"id,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	SizeBytes float64 `json:"sizeBytes,omitempty"`
	ThumbnailsUrl string `json:"thumbnailsUrl,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
}

type IncomeV2Income struct {
	AttachmentOverQuota bool `json:"attachmentOverQuota,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	Currency string `json:"currency,omitempty"`
	CustomerCiNumber string `json:"customerCiNumber,omitempty"`
	CustomerCountry string `json:"customerCountry"`
	CustomerId string `json:"customerId,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	CustomerNumber int `json:"customerNumber,omitempty"`
	CustomerVatNumber string `json:"customerVatNumber,omitempty"`
	Date string `json:"date"`
	DeletedDate string `json:"deletedDate,omitempty"`
	File *IncomeV2File `json:"file,omitempty"`
	FileId string `json:"fileId,omitempty"`
	Id string `json:"id,omitempty"`
	IncomeTaxDeductionAmount float64 `json:"incomeTaxDeductionAmount,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Lines []IncomeV2Lines `json:"lines"`
	Number int `json:"number,omitempty"`
	Payments []IncomeV2Payments `json:"payments,omitempty"`
	TaxGroups []IncomeV2TaxGroups `json:"taxGroups,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	TotalGrossAmount float64 `json:"totalGrossAmount,omitempty"`
	TotalNetAmount float64 `json:"totalNetAmount,omitempty"`
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	UnpaidAmount float64 `json:"unpaidAmount,omitempty"`
}

type IncomeV2Lines struct {
	AssetDepreciation []IncomeV2AssetDepreciation `json:"assetDepreciation,omitempty"`
	AssetLifetime int `json:"assetLifetime,omitempty"`
	AssetSalvageValue float64 `json:"assetSalvageValue,omitempty"`
	CategoryId string `json:"categoryId,omitempty"`
	CategoryType string `json:"categoryType,omitempty"`
	Description string `json:"description,omitempty"`
	GrossAmount float64 `json:"grossAmount,omitempty"`
	IncomeTaxDeductionAmount float64 `json:"incomeTaxDeductionAmount,omitempty"`
	NetAmount float64 `json:"netAmount,omitempty"`
	PriceDisplayType string `json:"priceDisplayType,omitempty"`
	ProductOrService string `json:"productOrService,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxEnabled bool `json:"taxEnabled"`
	TaxRate float64 `json:"taxRate"`
}

type IncomeV2Payments struct {
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

type IncomeV2TaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}