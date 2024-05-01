package apilambda

type WhitelistRequest struct {
	OldURL     string `json:"oldurl"`
	NewURL     string `json:"newurl"`
	MerchantID string `json:"merchantid"`
}
