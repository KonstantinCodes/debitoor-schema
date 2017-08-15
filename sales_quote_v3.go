package debitoor

type SalesQuoteV3History struct {
	Accepted string `json:"accepted,omitempty"`
	Invoiced string `json:"invoiced,omitempty"`
}

type SalesQuoteV3IncomeTaxDeductionGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}

type SalesQuoteV3Lines struct {
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
	ProductName string `json:"productName"`
	ProductOrService string `json:"productOrService,omitempty"`
	ProductSku string `json:"productSku,omitempty"`
	Quantity float64 `json:"quantity"`
	RecargoTaxAmount float64 `json:"recargoTaxAmount,omitempty"`
	RecargoTaxRate float64 `json:"recargoTaxRate,omitempty"`
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxEnabled bool `json:"taxEnabled"`
	TaxRate float64 `json:"taxRate"`
	UnitGrossPrice float64 `json:"unitGrossPrice,omitempty"`
	UnitId string `json:"unitId,omitempty"`
	UnitName string `json:"unitName,omitempty"`
	UnitNetPrice float64 `json:"unitNetPrice,omitempty"`
}

type SalesQuoteV3Links struct {
	Date string `json:"date,omitempty"`
	LinkId string `json:"linkId"`
	Type string `json:"type"`
}

type SalesQuoteV3Quote struct {
	Accepted bool `json:"accepted,omitempty"`
	AdditionalNotes string `json:"additionalNotes,omitempty"`
	Archived bool `json:"archived,omitempty"`
	BaseCurrency string `json:"baseCurrency,omitempty"`
	BaseCurrencyTotalGrossAmount float64 `json:"baseCurrencyTotalGrossAmount,omitempty"`
	BaseCurrencyTotalNetAmount float64 `json:"baseCurrencyTotalNetAmount,omitempty"`
	BaseCurrencyTotalNetDeductionAmount float64 `json:"baseCurrencyTotalNetDeductionAmount,omitempty"`
	BaseCurrencyTotalPensionFundAmount float64 `json:"baseCurrencyTotalPensionFundAmount,omitempty"`
	BaseCurrencyTotalRecargoTaxAmount float64 `json:"baseCurrencyTotalRecargoTaxAmount,omitempty"`
	BaseCurrencyTotalTaxAmount float64 `json:"baseCurrencyTotalTaxAmount,omitempty"`
	Booked bool `json:"booked,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	Currency string `json:"currency,omitempty"`
	CurrencyRate float64 `json:"currencyRate,omitempty"`
	CustomPaymentTermsDays int `json:"customPaymentTermsDays,omitempty"`
	CustomerAddress string `json:"customerAddress,omitempty"`
	CustomerCiNumber string `json:"customerCiNumber,omitempty"`
	CustomerCountry string `json:"customerCountry,omitempty"`
	CustomerCountryName string `json:"customerCountryName,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	CustomerVatNumber string `json:"customerVatNumber,omitempty"`
	Date string `json:"date,omitempty"`
	DeletedDate string `json:"deletedDate,omitempty"`
	DiscountRate float64 `json:"discountRate,omitempty"`
	DueDate string `json:"dueDate,omitempty"`
	History *SalesQuoteV3History `json:"history,omitempty"`
	Id string `json:"id,omitempty"`
	IncomeTaxDeductionGroups []SalesQuoteV3IncomeTaxDeductionGroups `json:"incomeTaxDeductionGroups,omitempty"`
	Invoiced bool `json:"invoiced,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Lines []SalesQuoteV3Lines `json:"lines"`
	Links []SalesQuoteV3Links `json:"links,omitempty"`
	Notes string `json:"notes,omitempty"`
	Number string `json:"number,omitempty"`
	PaymentTermsId int `json:"paymentTermsId,omitempty"`
	PensionFundRate float64 `json:"pensionFundRate,omitempty"`
	PensionFundType string `json:"pensionFundType,omitempty"`
	PriceDisplayType string `json:"priceDisplayType,omitempty"`
	RecargoTaxEnabled bool `json:"recargoTaxEnabled,omitempty"`
	RecargoTaxGroups []SalesQuoteV3RecargoTaxGroups `json:"recargoTaxGroups,omitempty"`
	SendDetails []SalesQuoteV3SendDetails `json:"sendDetails,omitempty"`
	Sent bool `json:"sent,omitempty"`
	State string `json:"state,omitempty"`
	TaxGroups []SalesQuoteV3TaxGroups `json:"taxGroups,omitempty"`
	Title string `json:"title,omitempty"`
	TotalGrossAmount float64 `json:"totalGrossAmount,omitempty"`
	TotalNetAmount float64 `json:"totalNetAmount,omitempty"`
	TotalNetAmountBeforeDiscount float64 `json:"totalNetAmountBeforeDiscount,omitempty"`
	TotalNetDeductionAmount float64 `json:"totalNetDeductionAmount,omitempty"`
	TotalNetDiscountAmount float64 `json:"totalNetDiscountAmount,omitempty"`
	TotalNetLineDiscountAmount float64 `json:"totalNetLineDiscountAmount,omitempty"`
	TotalPensionFundAmount float64 `json:"totalPensionFundAmount,omitempty"`
	TotalRecargoTaxAmount float64 `json:"totalRecargoTaxAmount,omitempty"`
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	Viewed bool `json:"viewed,omitempty"`
}

type SalesQuoteV3RecargoTaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}

type SalesQuoteV3SendDetails struct {
	Id string `json:"id"`
	Time string `json:"time"`
	To string `json:"to"`
	Viewed []string `json:"viewed"`
}

type SalesQuoteV3TaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}