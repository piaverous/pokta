package types

type OktaAuthResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_at"`
	Scope       string  `json:"scope"`
	TokenType   string  `json:"token_type"`
}
