package apiweb

type WhitelistRequest struct {
	OldURL string `json:"oldurl"`
	NewURL string `json:"newurl"`
	ID     string `json:"id"`
}
