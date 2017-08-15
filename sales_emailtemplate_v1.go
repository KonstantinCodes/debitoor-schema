package debitoor

type SalesEmailTemplateV1 struct {
	Id string `json:"id,omitempty"`
	Message string `json:"message"`
	Subject string `json:"subject"`
	Type string `json:"type"`
}