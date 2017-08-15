package debitoor

type LoginOauth2AuthorizeV1 struct {
	ClientId string `json:"client_id"`
	Lang string `json:"lang,omitempty"`
	RedirectUri string `json:"redirect_uri,omitempty"`
	ResponseType string `json:"response_type,omitempty"`
	State string `json:"state,omitempty"`
}