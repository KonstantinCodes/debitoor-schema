package debitoor

type DraftInvoice struct {
	AdditionalNotes string `json:"additionalNotes,omitempty"`
	BaseCurrency string `json:"baseCurrency,omitempty"`
	BaseCurrencyTotalGrossAmount float64 `json:"baseCurrencyTotalGrossAmount,omitempty"`
	BaseCurrencyTotalNetAmount float64 `json:"baseCurrencyTotalNetAmount,omitempty"`
	BaseCurrencyTotalNetDeductionAmount float64 `json:"baseCurrencyTotalNetDeductionAmount,omitempty"`
	BaseCurrencyTotalPensionFundAmount float64 `json:"baseCurrencyTotalPensionFundAmount,omitempty"`
	BaseCurrencyTotalRecargoTaxAmount float64 `json:"baseCurrencyTotalRecargoTaxAmount,omitempty"`
	BaseCurrencyTotalTaxAmount float64 `json:"baseCurrencyTotalTaxAmount,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	CreditedInvoiceId string `json:"creditedInvoiceId,omitempty"`
	Currency string `json:"currency,omitempty"`
	CurrencyRate float64 `json:"currencyRate,omitempty"`
	CustomPaymentTermsDays int `json:"customPaymentTermsDays,omitempty"`
	CustomerAddress string `json:"customerAddress,omitempty"`
	CustomerCiNumber string `json:"customerCiNumber,omitempty"`
	CustomerCountry string `json:"customerCountry,omitempty"`
	CustomerCountryName string `json:"customerCountryName,omitempty"`
	CustomerEmail string `json:"customerEmail,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	CustomerNumber int `json:"customerNumber,omitempty"`
	CustomerVatNumber string `json:"customerVatNumber,omitempty"`
	Date string `json:"date,omitempty"`
	DeletedDate string `json:"deletedDate,omitempty"`
	DiscountRate float64 `json:"discountRate,omitempty"`
	DisplayAsPaid bool `json:"displayAsPaid,omitempty"`
	DueDate string `json:"dueDate,omitempty"`
	History *History `json:"history,omitempty"`
	Id string `json:"id,omitempty"`
	IncomeTaxDeductionGroups []IncomeTaxDeductionGroups `json:"incomeTaxDeductionGroups,omitempty"`
	InvoiceTitle string `json:"invoiceTitle,omitempty"`
	InvoicedDeliveryNoteId string `json:"invoicedDeliveryNoteId,omitempty"`
	InvoicedQuoteId string `json:"invoicedQuoteId,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Lines []Lines `json:"lines,omitempty"`
	Links []Links `json:"links,omitempty"`
	Notes string `json:"notes,omitempty"`
	Number string `json:"number,omitempty"`
	OnlinePaymentProvider string `json:"onlinePaymentProvider,omitempty"`
	PaymentTermsId int `json:"paymentTermsId,omitempty"`
	PensionFundRate float64 `json:"pensionFundRate,omitempty"`
	PensionFundType string `json:"pensionFundType,omitempty"`
	PriceDisplayType string `json:"priceDisplayType,omitempty"`
	RecargoTaxEnabled bool `json:"recargoTaxEnabled,omitempty"`
	RecargoTaxGroups []RecargoTaxGroups `json:"recargoTaxGroups,omitempty"`
	SendDetails []SendDetails `json:"sendDetails,omitempty"`
	Sent bool `json:"sent,omitempty"`
	TaxGroups []TaxGroups `json:"taxGroups,omitempty"`
	TotalGrossAmount float64 `json:"totalGrossAmount,omitempty"`
	TotalNetAmount float64 `json:"totalNetAmount,omitempty"`
	TotalNetAmountBeforeDiscount float64 `json:"totalNetAmountBeforeDiscount,omitempty"`
	TotalNetDeductionAmount float64 `json:"totalNetDeductionAmount,omitempty"`
	TotalNetDiscountAmount float64 `json:"totalNetDiscountAmount,omitempty"`
	TotalNetLineDiscountAmount float64 `json:"totalNetLineDiscountAmount,omitempty"`
	TotalPensionFundAmount float64 `json:"totalPensionFundAmount,omitempty"`
	TotalRecargoTaxAmount float64 `json:"totalRecargoTaxAmount,omitempty"`
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	Type string `json:"type,omitempty"`
	Viewed bool `json:"viewed,omitempty"`
}

type History struct {
	Booked string `json:"booked,omitempty"`
}

type IncomeTaxDeductionGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
}

type Lines struct {
	BaseCurrencyGrossAmount float64 `json:"baseCurrencyGrossAmount,omitempty"`
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyPensionFundAmount float64 `json:"baseCurrencyPensionFundAmount,omitempty"`
	BaseCurrencyRecargoTaxAmount float64 `json:"baseCurrencyRecargoTaxAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	BaseCurrencyUnitGrossPrice float64 `json:"baseCurrencyUnitGrossPrice,omitempty"`
	BaseCurrencyUnitNetPrice float64 `json:"baseCurrencyUnitNetPrice,omitempty"`
	Description string `json:"description,omitempty"`
	GrossAmount float64 `json:"grossAmount,omitempty"`
	GrossAmountBeforeDiscount float64 `json:"grossAmountBeforeDiscount,omitempty"`
	IncomeTaxDeductionRate float64 `json:"incomeTaxDeductionRate,omitempty"`
	LineDiscountRate float64 `json:"lineDiscountRate,omitempty"`
	NetAmount float64 `json:"netAmount,omitempty"`
	NetAmountBeforeDiscount float64 `json:"netAmountBeforeDiscount,omitempty"`
	PensionFundAmount float64 `json:"pensionFundAmount,omitempty"`
	ProductId string `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
	ProductOrService string `json:"productOrService,omitempty"`
	ProductSku string `json:"productSku,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	RecargoTaxAmount float64 `json:"recargoTaxAmount,omitempty"`
	RecargoTaxRate float64 `json:"recargoTaxRate,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxEnabled bool `json:"taxEnabled,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
	UnitGrossPrice float64 `json:"unitGrossPrice,omitempty"`
	UnitId string `json:"unitId,omitempty"`
	UnitName string `json:"unitName,omitempty"`
	UnitNetPrice float64 `json:"unitNetPrice,omitempty"`
}

type Links struct {
	CreatedDate string `json:"createdDate,omitempty"`
	Date string `json:"date,omitempty"`
	LinkId string `json:"linkId,omitempty"`
	Type string `json:"type,omitempty"`
}

type RecargoTaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
}

type SendDetails struct {
	Id string `json:"id,omitempty"`
	Time string `json:"time,omitempty"`
	To string `json:"to,omitempty"`
	Viewed []string `json:"viewed,omitempty"`
}

type TaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
}
