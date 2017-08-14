package debitoor

type SalesDeliverynoteV2CompanyProfile struct {
	AccountHolderName string `json:"accountHolderName,omitempty"`
	Address string `json:"address,omitempty"`
	ApplyDeduction bool `json:"applyDeduction,omitempty"`
	BankAccount string `json:"bankAccount,omitempty"`
	BankAccount2 string `json:"bankAccount2,omitempty"`
	BankName string `json:"bankName,omitempty"`
	BankNumber string `json:"bankNumber,omitempty"`
	BankgirotNumber string `json:"bankgirotNumber,omitempty"`
	BicCode string `json:"bicCode,omitempty"`
	CashAccounting bool `json:"cashAccounting"`
	CommercialRegister string `json:"commercialRegister,omitempty"`
	CompanyNumber string `json:"companyNumber,omitempty"`
	CompanyType string `json:"companyType,omitempty"`
	Country string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	DistrictCourt string `json:"districtCourt,omitempty"`
	Email string `json:"email,omitempty"`
	EuVatNumber string `json:"euVatNumber,omitempty"`
	IbanCode string `json:"ibanCode,omitempty"`
	IncomeTaxDeduction float64 `json:"incomeTaxDeduction,omitempty"`
	Industry string `json:"industry,omitempty"`
	LogoUrl string `json:"logoUrl,omitempty"`
	Name string `json:"name,omitempty"`
	PensionFundRate float64 `json:"pensionFundRate,omitempty"`
	PensionFundType string `json:"pensionFundType,omitempty"`
	PlusgirotNumber string `json:"plusgirotNumber,omitempty"`
	ResponsibleName string `json:"responsibleName,omitempty"`
	ResponsiblePosition string `json:"responsiblePosition,omitempty"`
	SwiftCode string `json:"swiftCode,omitempty"`
	TaxEnabled bool `json:"taxEnabled"`
	TaxStatus string `json:"taxStatus,omitempty"`
	TelephoneNumber string `json:"telephoneNumber,omitempty"`
	VatNumber string `json:"vatNumber,omitempty"`
	WebSite string `json:"webSite,omitempty"`
}

type SalesDeliverynoteV2DeliveryNote struct {
	AdditionalNotes string `json:"additionalNotes,omitempty"`
	Archived bool `json:"archived,omitempty"`
	ArchivedDate string `json:"archivedDate,omitempty"`
	BaseCurrency string `json:"baseCurrency,omitempty"`
	BaseCurrencyTotalGrossAmount float64 `json:"baseCurrencyTotalGrossAmount,omitempty"`
	BaseCurrencyTotalNetAmount float64 `json:"baseCurrencyTotalNetAmount,omitempty"`
	BaseCurrencyTotalNetDeductionAmount float64 `json:"baseCurrencyTotalNetDeductionAmount,omitempty"`
	BaseCurrencyTotalPensionFundAmount float64 `json:"baseCurrencyTotalPensionFundAmount,omitempty"`
	BaseCurrencyTotalRecargoTaxAmount float64 `json:"baseCurrencyTotalRecargoTaxAmount,omitempty"`
	BaseCurrencyTotalTaxAmount float64 `json:"baseCurrencyTotalTaxAmount,omitempty"`
	BaseCurrencyUnpaidAmount float64 `json:"baseCurrencyUnpaidAmount,omitempty"`
	Booked bool `json:"booked,omitempty"`
	CompanyProfile *SalesDeliverynoteV2CompanyProfile `json:"companyProfile,omitempty"`
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
	Deleted bool `json:"deleted,omitempty"`
	DeletedByInvoiceId string `json:"deletedByInvoiceId,omitempty"`
	DeletedDate string `json:"deletedDate,omitempty"`
	DiscountRate float64 `json:"discountRate,omitempty"`
	DisplayAsPaid bool `json:"displayAsPaid,omitempty"`
	DueDate string `json:"dueDate,omitempty"`
	History *SalesDeliverynoteV2History `json:"history,omitempty"`
	Id string `json:"id,omitempty"`
	IncomeTaxDeductionGroups []SalesDeliverynoteV2IncomeTaxDeductionGroups `json:"incomeTaxDeductionGroups,omitempty"`
	InvoiceId string `json:"invoiceId,omitempty"`
	InvoiceTitle string `json:"invoiceTitle,omitempty"`
	InvoicedDeliveryNoteId string `json:"invoicedDeliveryNoteId,omitempty"`
	InvoicedQuoteId string `json:"invoicedQuoteId,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Lines []SalesDeliverynoteV2Lines `json:"lines"`
	Links []SalesDeliverynoteV2Links `json:"links,omitempty"`
	Notes string `json:"notes,omitempty"`
	Number string `json:"number,omitempty"`
	OnlinePaymentProvider string `json:"onlinePaymentProvider,omitempty"`
	OrderNumber float64 `json:"orderNumber,omitempty"`
	Paid bool `json:"paid,omitempty"`
	PaymentReceipts []SalesDeliverynoteV2PaymentReceipts `json:"paymentReceipts,omitempty"`
	PaymentTermsId float64 `json:"paymentTermsId,omitempty"`
	Payments []SalesDeliverynoteV2Payments `json:"payments,omitempty"`
	PensionFundRate float64 `json:"pensionFundRate,omitempty"`
	PensionFundType string `json:"pensionFundType,omitempty"`
	PriceDisplayType string `json:"priceDisplayType,omitempty"`
	QuoteId string `json:"quoteId,omitempty"`
	RecargoTaxEnabled bool `json:"recargoTaxEnabled,omitempty"`
	RecargoTaxGroups []SalesDeliverynoteV2RecargoTaxGroups `json:"recargoTaxGroups,omitempty"`
	SendDetails []SalesDeliverynoteV2SendDetails `json:"sendDetails,omitempty"`
	Sent bool `json:"sent,omitempty"`
	TaxGroups []SalesDeliverynoteV2TaxGroups `json:"taxGroups,omitempty"`
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
	UnpaidAmount float64 `json:"unpaidAmount,omitempty"`
	Viewed bool `json:"viewed,omitempty"`
}

type SalesDeliverynoteV2History struct {
	Booked string `json:"booked,omitempty"`
}

type SalesDeliverynoteV2IncomeTaxDeductionGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}

type SalesDeliverynoteV2Lines struct {
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

type SalesDeliverynoteV2Links struct {
	Amount float64 `json:"amount,omitempty"`
	CreatedDate string `json:"createdDate,omitempty"`
	Currency string `json:"currency,omitempty"`
	Date string `json:"date,omitempty"`
	LinkId string `json:"linkId"`
	Number string `json:"number,omitempty"`
	Type string `json:"type"`
	UnpaidAmountChange float64 `json:"unpaidAmountChange,omitempty"`
}

type SalesDeliverynoteV2PaymentReceipts struct {
}

type SalesDeliverynoteV2Payments struct {
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

type SalesDeliverynoteV2RecargoTaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}

type SalesDeliverynoteV2SendDetails struct {
	Id string `json:"id"`
	Time string `json:"time"`
	To string `json:"to"`
	Viewed []string `json:"viewed"`
}

type SalesDeliverynoteV2TaxGroups struct {
	BaseCurrencyNetAmount float64 `json:"baseCurrencyNetAmount,omitempty"`
	BaseCurrencyTaxAmount float64 `json:"baseCurrencyTaxAmount,omitempty"`
	Name string `json:"name,omitempty"`
	NetAmount float64 `json:"netAmount"`
	TaxAmount float64 `json:"taxAmount"`
	TaxRate float64 `json:"taxRate"`
}