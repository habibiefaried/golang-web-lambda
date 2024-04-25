package networkfirewall

type RequestBody struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
}
