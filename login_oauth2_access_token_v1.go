package debitoor

type LoginOauth2AccessTokenV1 struct {
	ClientSecret string `json:"client_secret"`
	Code string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}