package debitoor

type SalesDeliverynoteEmailV1 struct {
	AttachmentName string `json:"attachmentName,omitempty"`
	CcRecipient string `json:"ccRecipient,omitempty"`
	CopyMail bool `json:"copyMail,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Message string `json:"message,omitempty"`
	Recipient string `json:"recipient"`
	Subject string `json:"subject"`
}