package networkfirewallv2

// RequestBody will receive the ID and URL that needs to be whitelisted, so then, this must be processed
type RequestBody struct {
	// Data coming from http/kafka request
	ID  string `json:"id"`
	URL string `json:"url"`

	// After process
	IsTLS  bool
	Domain string
	Port   string
}
